package detector

import (
	"errors"
	"fmt"
	"github.com/bocchi-the-cache/stargazer/internal/entity"
	"github.com/bocchi-the-cache/stargazer/internal/model"
	"io/ioutil"
	"net/http"
	"net/url"
	"time"
)

type HttpDetector struct {
	BaseDetector
	HttpHost  string
	SslVerify bool
	SslExpire bool
}

func NewHttpDetector(task *entity.Task) (*HttpDetector, error) {
	tp := model.ProbeType(task.Type)
	if tp != model.HTTP && tp != model.HTTPS {
		return nil, errors.New(fmt.Sprintf("invalid probe type: %v", tp))
	}

	h := &HttpDetector{
		BaseDetector: BaseDetector{
			Name:       task.Name,
			Type:       model.ProbeType(task.Type),
			Target:     task.Target,
			Timeout:    time.Duration(task.Timeout) * time.Millisecond,
			Interval:   time.Duration(task.Interval) * time.Millisecond,
			UpdateTime: time.Now(),
			CreateTime: time.Now(),
		},
		HttpHost:  task.HttpHost,
		SslVerify: task.SslVerify,
		SslExpire: task.SslExpire,
	}

	return h, nil
}

func (h *HttpDetector) Detect() (entity.DataLog, error) {
	h.UpdateTime = time.Now()
	msg := newMessage()
	schema := string(h.Type)

	// make http request
	u := fmt.Sprintf("%s://%s", schema, h.Target)
	urlDetail, err := url.Parse(u)
	if err != nil {
		msg.Level = string(model.ERROR)
		msg.Message = fmt.Sprintf("make request error: %v", err)
		return msg, err
	}
	req, err := http.NewRequest("GET", u, nil)
	if err != nil {
		msg.Level = string(model.ERROR)
		msg.Message = fmt.Sprintf("make request error: %v", err)
		return msg, err
	}
	req.Header.Set("User-Agent", "Stargazer")
	// TODO: Host override in HTTP/HTTPS
	// Do http request
	client := &http.Client{
		Timeout: h.Timeout,
	}
	start := time.Now()
	resp, err := client.Do(req)
	if err != nil {
		msg.Level = string(model.ERROR)
		msg.Message = fmt.Sprintf("request error: %v", err)
		return msg, err
	}
	defer resp.Body.Close()
	// Read response body
	_, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		msg.Level = string(model.ERROR)
		msg.Message = fmt.Sprintf("read response body error: %v", err)
		return msg, err
	}
	// Check HTTP status code
	if resp.StatusCode != 200 {
		msg.Level = string(model.ERROR)
		msg.Message = fmt.Sprintf("http status code: %v", resp.StatusCode)
		return msg, nil
	}
	// Calculate response time
	end := time.Now()
	duration := end.Sub(start)
	if duration > h.Timeout {
		msg.Level = string(model.WARN)
		msg.Message = fmt.Sprintf("request timeout: %v", duration)
		return msg, nil
	}
	// Check SSL
	if schema == "https" {
		if h.SslVerify {
			if len(resp.TLS.PeerCertificates) == 0 {
				// defence: should never happen
				msg.Level = string(model.ERROR)
				msg.Message = "ssl verify failed: no ssl certificate found in resp"
				return msg, nil
			}
			if err := resp.TLS.PeerCertificates[0].VerifyHostname(urlDetail.Hostname()); err != nil {
				msg.Level = string(model.ERROR)
				msg.Message = fmt.Sprintf("ssl verify failed: %v", h.Target)
				return msg, nil
			}
		}
		if h.SslExpire {
			if resp.TLS.PeerCertificates[0].NotAfter.Before(time.Now().Add(7 * 24 * time.Hour)) {
				msg.Level = string(model.ERROR)
				msg.Message = fmt.Sprintf("ssl will expired: %v", h.Target)
				return msg, nil
			}
		}
	}

	msg.Message = fmt.Sprintf("request success. dest:%v, content-length:%v, "+
		"duration: %v", h.Target, resp.ContentLength, duration)
	return msg, nil
}
