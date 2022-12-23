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
	"github.com/sptuan/stargazer/internal/dao"
	"github.com/sptuan/stargazer/internal/entity"
	"github.com/sptuan/stargazer/pkg/logger"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

// AddTask - Add a new task
func AddTask(c *gin.Context) {
	var task Task
	if err := c.ShouldBindJSON(&task); err != nil {
		logger.Errorf("AddTask: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	t := &entity.Task{
		Name:        task.Name,
		Description: task.Description,
		Type:        task.Type,
		Target:      task.Target,
		HttpHost:    task.HttpHost,
		SslVerify:   task.SslVerify,
		SslExpire:   task.SslExpire,
		Interval:    task.Interval,
		Timeout:     task.Timeout,
	}

	if err := entity.Validate(t); err != nil {
		logger.Errorf("AddTask: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// TODO: add detector try_once here
	// if err := detector.TryOnce(t) ...

	if err := dao.AddTask(t); err != nil {
		logger.Errorf("AddTask: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	logger.Debugf("AddTask success: %v", t)
	c.JSON(http.StatusOK, gin.H{})
}

// DeleteTaskById - Delete task by ID
func DeleteTaskById(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(strings.TrimSpace(idStr))
	if err != nil {
		logger.Errorf("DeleteTaskById: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := dao.DeleteTaskById(id); err != nil {
		logger.Errorf("DeleteTaskById: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{})
}

// GetTaskById - Get task by ID
func GetTaskById(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(strings.TrimSpace(idStr))
	if err != nil {
		logger.Errorf("GetTaskById: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	task, err := dao.GetTaskById(id)
	if err != nil {
		logger.Errorf("GetTaskById: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"task": task})
}

// GetTasks - Get all tasks
func GetTasks(c *gin.Context) {
	tasks, err := dao.GetTasks()
	if err != nil {
		logger.Errorf("GetTasks: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"tasks": tasks})
}

// UpdateTaskById - Update task by ID
func UpdateTaskById(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(strings.TrimSpace(idStr))
	if err != nil {
		logger.Errorf("UpdateTaskById: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var task Task
	if err := c.ShouldBindJSON(&task); err != nil {
		logger.Errorf("UpdateTaskById: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	t := &entity.Task{
		Name:        task.Name,
		Description: task.Description,
		Type:        task.Type,
		Target:      task.Target,
		HttpHost:    task.HttpHost,
		SslVerify:   task.SslVerify,
		SslExpire:   task.SslExpire,
		Interval:    task.Interval,
		Timeout:     task.Timeout,
	}

	if err := entity.Validate(t); err != nil {
		logger.Errorf("UpdateTaskById: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := dao.UpdateTaskById(id, t); err != nil {
		logger.Errorf("UpdateTaskById: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{})
}