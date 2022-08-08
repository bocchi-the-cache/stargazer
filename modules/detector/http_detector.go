package detector

import (
	"errors"
	"fmt"
	"github.com/google/uuid"
	"github.com/sptuan/stargazer/modules/logger"
	"github.com/sptuan/stargazer/modules/model"
	"io/ioutil"
	"math"
	"net/http"
	"time"
)

const (
	SCHEME_HTTP  = "http"
	SCHEME_HTTPS = "https"

	METRIC_SSL_EXPIRE_AFTER = "ssl_expire_after"
	METRIC_RESPONSE_TIME    = "response_time"
	METRIC_STATUS_CODE      = "status_code"
)

type HttpDetector struct {
	BaseDetector
	Scheme  string
	Timeout time.Duration
	Client  *http.Client
}

func NewHttpDetector(name, target, scheme string, timeout time.Duration, interval time.Duration) (*HttpDetector, error) {
	u, err := uuid.NewUUID()
	if err != nil {
		return nil, err
	}
	h := &HttpDetector{
		BaseDetector: BaseDetector{
			Uuid:       u,
			Name:       name,
			Type:       DETECTOR_HTTP,
			Target:     target,
			Interval:   interval,
			UpdateTime: time.Now(),
			CreateTime: time.Now(),
		},
		Scheme:  scheme,
		Timeout: timeout,
	}
	h.Init()
	return h, nil
}

func (h *HttpDetector) Init() {
	h.Signal = make(chan struct{})
	// use mostly default client config
	h.Client = &http.Client{
		Timeout: h.Timeout,
	}
}

func (h *HttpDetector) Start() error {
	go func() {
		h.StartInterval()
	}()
	return nil
}

func (h *HttpDetector) Stop() error {
	h.Signal <- struct{}{}
	return nil
}

func (h *HttpDetector) StartInterval() {
	for {
		select {
		case <-h.Signal:
			logger.Infof("detector %v stopped", h)
			return
		case <-time.After(h.Interval):
			h.UpdateTime = time.Now()
			h.DetectOnce()
		}
	}
}

func (h *HttpDetector) DetectOnce() {
	// TODO: Report Key Error to collector
	msg := h.NewMessage()
	// Do Key Detection
	err := h.DetectHTTP(&msg)
	if err != nil {
		logger.Debugf("detect http failed. detector: %v, err: %s", h, err)
	}
	err = h.Validate(&msg)
	if err != nil {
		logger.Debugf("validate http failed. detector: %v, err: %s", h, err)

	}
	logger.Debugf("detect http success. detector: %v, duration: %s", h, msg.Report.Metric[METRIC_RESPONSE_TIME])
}

// DetectHTTP do key detection, return a message with health report
func (h *HttpDetector) DetectHTTP(msg *model.DetectorMessage) error {
	t1 := time.Now()

	// make request
	req, err := http.NewRequest(http.MethodGet, h.Scheme+"://"+h.Target, nil)
	if err != nil {
		msg.Report.Metric[METRIC_RESPONSE_TIME] = time.Since(t1).Seconds()
		msg.Report.Errs = append(msg.Report.Errs, err)
		return err
	}

	// do request
	resp, err := h.Client.Do(req)
	if err != nil {
		msg.Report.Metric[METRIC_RESPONSE_TIME] = time.Since(t1).Seconds()
		msg.Report.Errs = append(msg.Report.Errs, err)
		return err
	}
	msg.Report.Metric[METRIC_STATUS_CODE] = float64(resp.StatusCode)

	// read response
	defer resp.Body.Close()
	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		msg.Report.Metric[METRIC_RESPONSE_TIME] = time.Since(t1).Seconds()
		msg.Report.Errs = append(msg.Report.Errs, err)
		return err
	}
	logger.Debugf("fetch http body success. detector: %v, body length: %s", h, len(respBody))

	// Check status code
	if resp.StatusCode != http.StatusOK {
		msg.Report.Metric[METRIC_RESPONSE_TIME] = time.Since(t1).Seconds()
		msg.Report.Errs = append(msg.Report.Errs, err)
		return errors.New(fmt.Sprintf("status code not 200, code: %v", resp.StatusCode))
	}

	// Check SSL expire time
	if h.Scheme == SCHEME_HTTPS {
		minExpireAfter := math.MaxFloat64
		for _, cert := range resp.TLS.PeerCertificates {
			if cert.NotAfter.Sub(time.Now()).Seconds() < minExpireAfter {
				minExpireAfter = cert.NotAfter.Sub(time.Now()).Seconds()
			}
			msg.Report.Metric[METRIC_SSL_EXPIRE_AFTER] = minExpireAfter
			if cert.NotAfter.Before(time.Now()) {
				msg.Report.Metric[METRIC_RESPONSE_TIME] = time.Since(t1).Seconds()
				msg.Report.Errs = append(msg.Report.Errs, err)
				return errors.New(fmt.Sprintf("ssl expire time: %s", cert.NotAfter))
			}
		}
	}
	msg.Report.Metric[METRIC_RESPONSE_TIME] = time.Since(t1).Seconds()
	return nil
}
