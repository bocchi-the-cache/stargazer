package detector

import "net/http"

const (
	SCHEME_HTTP  = "http"
	SCHEME_HTTPS = "https"
)

type HttpDetector struct {
	BaseDetector
	Scheme string
	Client http.Client
}
