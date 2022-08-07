package detector

import (
	"github.com/google/uuid"
	"time"
)

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
}

type Detector interface {
	Init() error
	Start() error
	Stop() error
}
