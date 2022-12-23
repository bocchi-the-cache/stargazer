package dao

import (
	"github.com/sptuan/stargazer/internal/constant"
	"github.com/sptuan/stargazer/internal/db"
	"github.com/sptuan/stargazer/internal/entity"
	"time"
)

func AddDataLog(dataLog *entity.DataLog) error {
	err := db.Db.Create(dataLog).Error
	return err
}

func GetDataLogByTaskIdInTimeRange(taskId int, startTime time.Time, endTime time.Time) ([]entity.DataLog, error) {
	var dataLogs []entity.DataLog
	err := db.Db.Where("task_id = ? AND created_at BETWEEN ? AND ?", taskId, startTime, endTime).Find(&dataLogs).Error
	return dataLogs, err
}

func GetDataLogByTaskIdLevelInTimeRange(taskId int, level constant.Level, startTime time.Time, endTime time.Time) ([]entity.DataLog, error) {
	var dataLogs []entity.DataLog
	err := db.Db.Where("task_id = ? AND level = ? AND created_at BETWEEN ? AND ?", taskId, level, startTime, endTime).Find(&dataLogs).Error
	return dataLogs, err
}
