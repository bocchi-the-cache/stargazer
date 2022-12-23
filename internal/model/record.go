package model

import "time"

type UserDetectionRecord struct {
	Id          int       `json:"id" form:"id" gorm:"primaryKey;autoIncrement"`
	Timestamp   time.Time `json:"timestamp" form:"timestamp" gorm:"index:idx_name_timestamp,priority:10"`
	Uuid        string    `json:"uuid" form:"uuid"`
	Name        string    `json:"name" form:"name" gorm:"index:idx_name_timestamp,priority:5"`
	Type        string    `json:"type" form:"type"`
	Target      string    `json:"target" form:"target"`
	HealthLevel int       `json:"health_level" form:"health_level"`
	Errs        string    `json:"errs" form:"errs"`
	Metric      string    `json:"metric" form:"metric"`
}
