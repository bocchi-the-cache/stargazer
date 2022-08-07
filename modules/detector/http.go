package detector

import (
	"errors"
	"fmt"
	"github.com/sptuan/stargazer/modules/logger"
	"io/ioutil"
	"net/http"
	"time"
)

const (
	SCHEME_HTTP  = "http"
	SCHEME_HTTPS = "https"
)

type HttpDetector struct {
	BaseDetector
	Scheme  string
	Timeout time.Duration
	Client  *http.Client
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
			return
		case <-time.After(h.Interval):
			h.DetectOnce()
		}
	}
}

func (h *HttpDetector) DetectOnce() {
	t1 := time.Now()
	err := h.DetectHTTP()
	t := time.Now().After(t1)

	if err != nil {
		// TODO: Report Error to collector
		logger.Errorf("detect http failed. detector: %v, err: %s", h, err)
	}
	// TODO: Report Success to collector
	logger.Infof("detect http success. detector: %v, duration: %s", h, t)
}

func (h *HttpDetector) DetectHTTP() error {
	req, err := http.NewRequest(http.MethodGet, h.Scheme+h.Target, nil)
	if err != nil {
		return err
	}
	resp, err := h.Client.Do(req)
	if err != nil {
		return err
	}

	defer resp.Body.Close()
	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	logger.Infof("detect http success. detector: %v, body: %s", h, string(respBody))

	// Check status code
	if resp.StatusCode != http.StatusOK {
		return errors.New(fmt.Sprintf("status code not 200, code: %i", resp.StatusCode))
	}

	// Check SSL expire time
	if h.Scheme == SCHEME_HTTPS {
		for _, cert := range resp.TLS.PeerCertificates {
			if cert.NotAfter.Before(time.Now()) {
				return errors.New(fmt.Sprintf("ssl expire time: %s", cert.NotAfter))
			}
		}
	}

	return nil
}
