package database

import (
	"github.com/sptuan/stargazer/modules/logger"
	"github.com/sptuan/stargazer/modules/model"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"io/fs"
	"os"
	"path"
)

// NOTE: For easy use and small projects, we use gorm.DB (sqlite3).
// Actually, we should use time-series database like prometheus or influxdb.
// In our roadmap, we provide metrics exporter for prometheus to fetch.
var db *gorm.DB

func Init(dbPath string) error {
	dir := path.Dir(dbPath)
	err := os.MkdirAll(dir, fs.ModeDir)
	if err != nil {
		return err
	}
	// TODO: Debug/Production log mode
	db, err = gorm.Open(sqlite.Open(dbPath), &gorm.Config{})
	if err != nil {
		logger.Panicf("failed to open database: %v", err)
		return err
	}

	err = initDetectionRecord()
	if err != nil {
		return err
	}
	return nil
}

func initDetectionRecord() error {
	return db.AutoMigrate(&model.UserDetectionRecord{})
}
