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

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetDataLogByTaskId - Get data log by task ID
func GetDataLogByTaskId(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}

// GetDataSeriesByTaskId - Get data series by task ID
func GetDataSeriesByTaskId(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}

// GetDataStatusByTaskId - Get data status by task ID
func GetDataStatusByTaskId(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}