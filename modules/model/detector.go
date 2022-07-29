package model

import (
	"github.com/google/uuid"
	"time"
)

type HealthReport struct {
	HealthLevel
	Errs   []error
	Metric map[string]float64
	Logs   []string
}

type DetectorMessage struct {
	// uuid is unique, mostly for index
	Uuid       uuid.UUID
	Name       string
	Type       string
	Target     string
	UpdateTime time.Time
}
