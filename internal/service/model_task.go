/*
 * Stargazer
 *
 * Stargazer Backend OpenAPI Specification
 *
 * API version: 0.1.0
 * Contact: sptuan@steinslab.io
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package service

type Task struct {
	Id int64 `json:"id,omitempty"`

	Name string `json:"name,omitempty"`

	Description string `json:"description,omitempty"`

	Type string `json:"type,omitempty"`

	// url/ip:port
	Target string `json:"target,omitempty"`

	// Overwrite target host (in http/https)
	HttpHost string `json:"http_host,omitempty"`

	// Verify SSL certificate (in http/https)
	SslVerify bool `json:"ssl_verify,omitempty"`

	// Check SSL certificate expiration soon (in http/https)
	SslExpire bool `json:"ssl_expire,omitempty"`

	// Interval in seconds
	Interval int64 `json:"interval,omitempty"`

	// Timeout in milliseconds
	Timeout int64 `json:"timeout,omitempty"`

	Status string `json:"status,omitempty"`

	CreatedAt int64 `json:"created_at,omitempty"`

	UpdatedAt int64 `json:"updated_at,omitempty"`
}
