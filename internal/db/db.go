package db

import (
	"github.com/sptuan/stargazer/internal/model"
	"github.com/sptuan/stargazer/pkg/logger"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"io/fs"
	"os"
	"path"
)

// NOTE: For easy use and small projects, we use gorm.DB (sqlite3).
// Should use time-series database like prometheus in the future.
var Db *gorm.DB

func Init(dbPath string) error {
	dir := path.Dir(dbPath)
	err := os.MkdirAll(dir, fs.ModeDir)
	if err != nil {
		return err
	}
	// TODO: Debug/Production log mode
	Db, err = gorm.Open(sqlite.Open(dbPath), &gorm.Config{})
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
	return Db.AutoMigrate(&model.UserDetectionRecord{})
}
