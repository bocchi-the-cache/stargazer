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

type DataStatus struct {
	Status string `json:"status,omitempty"`

	LastRun int64 `json:"last_run,omitempty"`

	NextRun int64 `json:"next_run,omitempty"`
}