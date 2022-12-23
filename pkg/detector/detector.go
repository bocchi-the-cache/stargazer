package detector

import (
	"github.com/google/uuid"
	"github.com/sptuan/stargazer/internal/model"
	"time"
)

const DETECTOR_HTTP = "http_detector"

type BaseDetector struct {
	// uuid is unique, mostly for index
	Uuid       uuid.UUID
	Name       string
	Type       string
	Target     string
	CreateTime time.Time
	UpdateTime time.Time
	Interval   time.Duration
	Signal     chan struct{}
	Validators []func(*model.DetectorMessage) error
}

type Detector interface {
	Start() error
	Stop() error
}

func (d *BaseDetector) NewMessage() model.DetectorMessage {
	return model.DetectorMessage{
		Uuid:       d.Uuid,
		Name:       d.Name,
		Type:       d.Type,
		Target:     d.Target,
		UpdateTime: time.Now(),
		Report: model.HealthReport{
			HealthLevel: model.LevelHealthy,
			Errs:        []error{},
			Metric:      map[string]float64{},
			Logs:        []string{},
		},
	}
}

// RegisterValidator registers a validator function
// NOTE: Validators only return error of itself.
// For alert error, put it into message->report->errs.
func (d *BaseDetector) RegisterValidator(fn func(msg *model.DetectorMessage) error) {
	d.Validators = append(d.Validators, fn)
}

func (d *BaseDetector) Validate(msg *model.DetectorMessage) error {
	for _, fn := range d.Validators {
		if err := fn(msg); err != nil {
			return err
		}
	}
	return nil
}
