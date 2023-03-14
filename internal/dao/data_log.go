package dao

import (
	"github.com/bocchi-the-cache/stargazer/internal/db"
	"github.com/bocchi-the-cache/stargazer/internal/entity"
	"github.com/bocchi-the-cache/stargazer/internal/model"
	"time"
)

func AddDataLog(dataLog *entity.DataLog) error {
	dataLog.CreatedAt = time.Now().Unix()
	err := db.Db.Create(dataLog).Error
	return err
}

func GetDataLogByTaskIdInTimeRange(taskId int, startTime int, endTime int) ([]entity.DataLog, error) {
	var dataLogs []entity.DataLog
	err := db.Db.Where("task_id = ? AND created_at BETWEEN ? AND ?", taskId, startTime, endTime).Find(&dataLogs).Error
	return dataLogs, err
}

func GetDataLogByTaskIdLevelInTimeRange(taskId int, level model.Level, startTime int, endTime int) ([]entity.DataLog, error) {
	var dataLogs []entity.DataLog
	err := db.Db.Where("task_id = ? AND level = ? AND created_at BETWEEN ? AND ?", taskId, level, startTime, endTime).Find(&dataLogs).Error
	return dataLogs, err
}

func GetDataLogLastByTaskId(taskId int) (*entity.DataLog, error) {
	var dataLog entity.DataLog
	err := db.Db.Where("task_id = ?", taskId).Last(&dataLog).Error
	return &dataLog, err
}

func DeleteDataLogBeforeTime(time int) error {
	err := db.Db.Where("created_at < ?", time).Delete(&entity.DataLog{}).Error
	return err
}
