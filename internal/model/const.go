package model

type Level string

type ProbeType string

const (
	DEBUG Level = "debug"
	INFO  Level = "info"
	WARN  Level = "warn"
	ERROR Level = "error"
)

const (
	PING  ProbeType = "ping"
	HTTP  ProbeType = "http"
	HTTPS ProbeType = "https"
	PORT  ProbeType = "port"
)
