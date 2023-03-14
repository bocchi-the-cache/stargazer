package dao

import (
	"github.com/bocchi-the-cache/stargazer/internal/db"
	"github.com/bocchi-the-cache/stargazer/internal/entity"
)

func AddTask(task *entity.Task) error {
	err := db.Db.Create(task).Error
	return err
}

func GetTaskById(id int) (*entity.Task, error) {
	var task entity.Task
	err := db.Db.First(&task, id).Error
	return &task, err
}

func GetTasks() ([]*entity.Task, error) {
	var tasks []*entity.Task
	err := db.Db.Find(&tasks).Error
	return tasks, err
}

func UpdateTaskById(id int, task *entity.Task) error {
	err := db.Db.Model(&task).Where("id = ?", id).Updates(task).Error
	return err
}

func DeleteTaskById(id int) error {
	err := db.Db.Delete(&entity.Task{}, id).Error
	return err
}
