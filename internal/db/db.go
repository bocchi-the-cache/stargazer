package db

import (
	"github.com/bocchi-the-cache/stargazer/internal/conf"
	"github.com/bocchi-the-cache/stargazer/internal/entity"
	"github.com/bocchi-the-cache/stargazer/internal/model"
	"github.com/bocchi-the-cache/stargazer/pkg/logger"
	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
	"io/fs"
	"os"
	"path"
)

// Db NOTE: For easy use and small projects, we use gorm.DB (sqlite3).
// Should use time-series database like prometheus in the future.
var Db *gorm.DB

func Init(dbPath string) error {
	dir := path.Dir(dbPath)
	err := os.MkdirAll(dir, fs.ModeDir)
	if err != nil {
		return err
	}
	Db, err = gorm.Open(sqlite.Open(dbPath), &gorm.Config{})
	if err != nil {
		logger.Panicf("failed to open database: %v", err)
		return err
	}

	sqlDB, err := Db.DB()
	if err != nil {
		logger.Panicf("failed to get sql db: %v", err)
		return err
	}

	sqlDB.SetMaxIdleConns(0)
	sqlDB.SetMaxOpenConns(1)

	if model.Level(conf.Cfg.Service.LogLevel) == model.DEBUG {
		Db = Db.Debug()
	}

	err = Db.AutoMigrate(&entity.Task{}, &entity.DataLog{})
	if err != nil {
		return err
	}
	return nil
}
