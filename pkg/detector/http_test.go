package detector

import (
	"github.com/sptuan/stargazer/internal/model"
	"github.com/sptuan/stargazer/pkg/logger"
	"testing"
	"time"
)

func initLogger() {
	logger.Init()
}

func TestHttpDetector(t *testing.T) {
	initLogger()

	d, err := NewHttpDetector(
		"test-detector",
		"steinslab.io",
		SCHEME_HTTPS,
		time.Second*10,
		time.Second*10)
	if err != nil {
		t.Errorf("failed to create http detector: %v", err)
	}
	d.RegisterValidatorResponseTime(time.Millisecond*1000, model.LevelCritical)
	d.RegisterValidatorSSLExpireAfter(time.Hour*24*7, model.LevelWarning)
	_ = d.Start()
	time.Sleep(time.Second * 60)
	_ = d.Stop()
	time.Sleep(time.Second * 60)
}
