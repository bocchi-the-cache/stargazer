package constant

type Level string

type ProbeType string

const (
	DebugLevel Level = "debug"
	INFO       Level = "info"
	WARN       Level = "warn"
	ERROR      Level = "error"
)

const (
	PING  ProbeType = "ping"
	HTTP  ProbeType = "http"
	HTTPS ProbeType = "https"
	PORT  ProbeType = "port"
)
