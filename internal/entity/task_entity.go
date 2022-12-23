package entity

type Task struct {
	Id          int    `json:"id" form:"id" gorm:"primaryKey;autoIncrement"`
	Name        string `json:"name" form:"name" gorm:"type:varchar(255);not null;uniqueIndex"`
	Description string `json:"description" form:"description" gorm:"type:text"`
	Type        string `json:"type" form:"type" gorm:"type:varchar(255);not null"`
	Target      string `json:"target" form:"target" gorm:"type:longtext;not null"`
	HttpHost    string `json:"http_host" form:"http_host" gorm:"type:varchar(255)"`
	SslVerify   bool   `json:"ssl_verify" form:"ssl_verify" gorm:"type:tinyint(1);not null;default:0"`
	SslExpire   bool   `json:"ssl_expire" form:"ssl_expire" gorm:"type:tinyint(1);not null;default:0"`
	Interval    int    `json:"interval" form:"interval" gorm:"type:int(11);not null;default:60"`
	Timeout     int    `json:"timeout" form:"timeout" gorm:"type:int(11);not null;default:10000"`
	Status      string `json:"status" form:"status" gorm:"type:varchar(255);not null;default:'active'"`
	CreatedAt   int    `json:"created_at" form:"created_at" gorm:"type:int(11);not null;default:0"`
	UpdatedAt   int    `json:"updated_at" form:"updated_at" gorm:"type:int(11);not null;default:0"`
}

func (Task) TableName() string {
	return "sgr_tasks"
}
