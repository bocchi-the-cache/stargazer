package entity

import (
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

type Task struct {
	gorm.Model
	Name        string `json:"name" form:"name" gorm:"type:varchar(255);not null;uniqueIndex" validate:"required"`
	Description string `json:"description" form:"description" gorm:"type:text"`
	Type        string `json:"type" form:"type" gorm:"type:varchar(255);not null" validate:"required,oneof=http https ping port"`
	Target      string `json:"target" form:"target" gorm:"type:longtext;not null" validate:"required"`
	HttpHost    string `json:"http_host" form:"http_host" gorm:"type:varchar(255)"`
	SslVerify   bool   `json:"ssl_verify" form:"ssl_verify" gorm:"type:tinyint(1);not null;default:0"`
	SslExpire   bool   `json:"ssl_expire" form:"ssl_expire" gorm:"type:tinyint(1);not null;default:0"`
	Interval    int64  `json:"interval" form:"interval" gorm:"type:int(11);not null;default:60" validate:"required,gte=0"`
	Timeout     int64  `json:"timeout" form:"timeout" gorm:"type:int(11);not null;default:10000" validate:"required,gte=0"`
	Status      string `json:"status" form:"status" gorm:"type:varchar(255);not null;default:'active'"`
}

func (Task) TableName() string {
	return "sgr_tasks"
}

func (e *Task) Validate() error {
	// base check based on tag
	validate := validator.New()
	if err := validate.Struct(e); err != nil {
		return err
	}

	// custom check
	return nil
}
