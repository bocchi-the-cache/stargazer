package detector

import (
	"errors"
	"fmt"
	"github.com/sptuan/stargazer/internal/model"
	"time"
)

func (h *HttpDetector) RegisterValidatorResponseTime(t time.Duration, lvl model.HealthLevel) {
	fn := func(msg *model.DetectorMessage) error {
		if msg == nil {
			return errors.New("RegisterValidatorResponseTime: msg is nil")
		}
		duration, ok := msg.Report.Metric[METRIC_RESPONSE_TIME]
		if !ok {
			return errors.New("RegisterValidatorResponseTime: metric response_time not found")
		}
		if duration > t.Seconds() {
			if msg.Report.HealthLevel < lvl {
				msg.Report.HealthLevel = lvl
			}
			msg.Report.Errs = append(msg.Report.Errs, errors.New(fmt.Sprintf("response time: %v s > %v", duration, t)))
		}
		return nil
	}
	h.RegisterValidator(fn)
}

func (h *HttpDetector) RegisterValidatorSSLExpireAfter(t time.Duration, lvl model.HealthLevel) {
	fn := func(msg *model.DetectorMessage) error {
		if msg == nil {
			return errors.New("RegisterValidatorSSLExpireAfter: msg is nil")
		}
		duration, ok := msg.Report.Metric[METRIC_SSL_EXPIRE_AFTER]
		if !ok {
			return errors.New("RegisterValidatorSSLExpireAfter: metric ssl_expire_after not found")
		}
		if duration < t.Seconds() {
			if msg.Report.HealthLevel < lvl {
				msg.Report.HealthLevel = lvl
			}
			msg.Report.Errs = append(msg.Report.Errs, errors.New(fmt.Sprintf("ssl expire time: %v s > %v", duration, t)))
		}
		return nil
	}
	h.RegisterValidator(fn)
}
