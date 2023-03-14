package db

import (
	"github.com/bocchi-the-cache/stargazer/internal/entity"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"testing"
)

func TestNewSqliteDb(t *testing.T) {
	// auto migrate sqlite db
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		t.Error(err)
	}
	err = db.AutoMigrate(&entity.Task{}, &entity.DataLog{})
	if err != nil {
		t.Error(err)
	}

	// insert data
	task := entity.Task{
		Name:        "test",
		Description: "test",
		Target:      "http://localhost:8080",
		Interval:    10,
		Timeout:     1000,
	}
	err = db.Create(&task).Error
	if err != nil {
		t.Error(err)
	}
}
