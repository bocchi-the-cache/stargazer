package db

import (
	"github.com/glebarez/sqlite"
	"github.com/sptuan/stargazer/internal/conf"
	"github.com/sptuan/stargazer/internal/entity"
	"github.com/sptuan/stargazer/internal/model"
	"github.com/sptuan/stargazer/pkg/logger"
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

	if model.Level(conf.Cfg.Service.LogLevel) == model.DEBUG {
		Db = Db.Debug()
	}

	err = Db.AutoMigrate(&entity.Task{}, &entity.DataLog{})
	if err != nil {
		return err
	}
	return nil
}
