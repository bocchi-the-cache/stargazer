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
	"fmt"
	"github.com/bocchi-the-cache/stargazer/internal/dao"
	"github.com/bocchi-the-cache/stargazer/internal/entity"
	"github.com/bocchi-the-cache/stargazer/internal/model"
	"github.com/bocchi-the-cache/stargazer/pkg/logger"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

// GetDataLogByTaskId - Get data log by task ID
func GetDataLogByTaskId(c *gin.Context) {
	// parse task id, start, end
	taskIdStr := c.Param("taskId")
	taskId, err := strconv.Atoi(taskIdStr)
	if err != nil {
		logger.Errorf("GetDataLogByTaskId: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	level := c.Query("level")

	var start, end int
	startStr := c.Query("start")
	if startStr != "" {
		start, err = strconv.Atoi(startStr)
		if err != nil {
			logger.Errorf("GetDataLogByTaskId: %v", err)
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
	} else {
		start = int(time.Now().Add(-1 * time.Hour).Unix())
	}

	endStr := c.Query("end")
	if endStr != "" {
		end, err = strconv.Atoi(endStr)
		if err != nil {
			logger.Errorf("GetDataLogByTaskId: %v", err)
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
	} else {
		end = int(time.Now().Unix())
	}

	dataLog := []entity.DataLog{}

	if level == "" {
		dataLog, err = dao.GetDataLogByTaskIdInTimeRange(taskId, start, end)
		if err != nil {
			logger.Errorf("GetDataLogByTaskId: %v", err)
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
	} else {
		dataLog, err = dao.GetDataLogByTaskIdLevelInTimeRange(taskId, model.Level(level), start, end)
		if err != nil {
			logger.Errorf("GetDataLogByTaskId: %v", err)
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
	}

	c.JSON(http.StatusOK, gin.H{"log": dataLog})
}

// parseDataLogQuery - parse query params for data log
func parseDataLogQuery(c *gin.Context) (int, int, int, int, error) {
	taskId, err := strconv.Atoi(c.Param("taskId"))
	if err != nil {
		return 0, 0, 0, 0, err
	}

	// parse interval
	intervalStr := c.Query("interval")
	interval := 0
	if intervalStr == "" {
		interval = 3600000
	} else {
		interval, err = strconv.Atoi(intervalStr)
		if err != nil {
			return 0, 0, 0, 0, err
		}
	}

	// parse start
	startStr := c.Query("start")
	start := 0
	if startStr == "" {
		start = int(time.Now().Add(-24 * time.Hour).Unix())
	} else {
		start, err = strconv.Atoi(startStr)
		if err != nil {
			return 0, 0, 0, 0, err
		}
	}

	endStr := c.Query("end")
	end := 0
	if endStr == "" {
		end = int(time.Now().Unix())
	} else {
		end, err = strconv.Atoi(endStr)
		if err != nil {
			return 0, 0, 0, 0, err
		}
	}
	// parse end
	return taskId, interval, start, end, nil
}

// GetDataSeriesByTaskId - Get data series by task ID
func GetDataSeriesByTaskId(c *gin.Context) {
	taskId, interval, start, end, err := parseDataLogQuery(c)
	if err != nil {
		logger.Errorf("GetDataSeriesByTaskId: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// get data log
	dataLog, err := dao.GetDataLogByTaskIdInTimeRange(taskId, start, end)
	var dataSeries []DataSeries
	if len(dataLog) == 0 {
		c.JSON(http.StatusOK, gin.H{"series": dataSeries})
		return
	}

	for st := start; st < end; st += interval {
		var dataSeriesItem DataSeries
		dataSeriesItem.TimeStart = int64(st)
		dataSeriesItem.TimeEnd = int64(st + interval)

		HealtyCounter := int64(0)
		UnhealtyCounter := int64(0)
		msg := strings.Builder{}
		for index := range dataLog {
			logTime := int(dataLog[index].CreatedAt)
			if logTime <= st || logTime > st+interval {
				continue
			}
			if model.INFO == model.Level(dataLog[index].Level) {
				HealtyCounter++
			} else {
				UnhealtyCounter++
				timeHuman := time.Unix(int64(dataLog[index].CreatedAt), 0).Format("2006-01-02 15:04:05")
				msg.WriteString(fmt.Sprintf("%s:%s, %s\n", timeHuman, dataLog[index].Level, dataLog[index].Message))
			}
		}
		if (HealtyCounter + UnhealtyCounter) == 0 {
			dataSeriesItem.Value = -1
		} else {
			dataSeriesItem.Value = float32(HealtyCounter) / float32(HealtyCounter+UnhealtyCounter)
		}
		dataSeriesItem.SuccessCount = HealtyCounter
		dataSeriesItem.FailCount = UnhealtyCounter
		dataSeries = append(dataSeries, dataSeriesItem)
	}

	c.JSON(http.StatusOK, gin.H{"series": dataSeries})
}

// GetDataStatusByTaskId - Get data status by task ID
func GetDataStatusByTaskId(c *gin.Context) {
	idStr := c.Param("taskId")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		logger.Errorf("GetDataStatusByTaskId: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	dataStatus, err := dao.GetDataLogLastByTaskId(id)
	if err != nil {
		logger.Errorf("GetDataStatusByTaskId: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": dataStatus})
}
