package model

type HealthLevel int

const (
	LevelHealthy HealthLevel = iota + 1
	LevelWarning
	LevelCritical
)
