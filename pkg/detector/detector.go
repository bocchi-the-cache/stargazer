package detector

import (
	"github.com/bocchi-the-cache/stargazer/internal/entity"
	"github.com/bocchi-the-cache/stargazer/internal/model"
	"time"
)

const DETECTOR_HTTP = "http_detector"

type BaseDetector struct {
	Id         int
	Name       string
	Type       model.ProbeType
	Target     string
	Timeout    time.Duration
	Interval   time.Duration
	CreateTime time.Time
	UpdateTime time.Time
}

type Detector interface {
	Detect() (entity.DataLog, error)
}

func newMessage() entity.DataLog {
	return entity.DataLog{
		CreatedAt: time.Now().Unix(),
		Level:     string(model.INFO),
	}
}
