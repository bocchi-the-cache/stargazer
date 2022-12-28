package entity

type DataLog struct {
	Id        int    `json:"id" form:"id" gorm:"primaryKey;autoIncrement"`
	TaskId    int    `json:"task_id" form:"task_id" gorm:"type:int(11);not null"`
	CreatedAt int64  `json:"created_at" form:"created_at" gorm:"type:int(11);not null"`
	Level     string `json:"level" form:"level" gorm:"type:varchar(255);not null"`
	Message   string `json:"message" form:"message" gorm:"type:text;not null"`
}

func (DataLog) TableName() string {
	return "sgr_data_logs"
}
