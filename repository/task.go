package repository

import (
	"main.go/model"
	"gorm.io/gorm"
)

type TaskRepository interface {
    Create(task *model.Task) (*model.Task, error)
    GetByID(id int) (*model.Task, error)
    Update(id int, task *model.Task) error
    Delete(id int) error
}

type taskRepository struct {
	db *gorm.DB
}

func NewTaskRepo(db *gorm.DB) *taskRepository {
	return &taskRepository{db}
}


func (t *taskRepository) Create(task *model.Task) (*model.Task, error) {
	if err := t.db.Create(task).Error; err != nil {
		return nil, err
	}
	return task, nil
}

func (t *taskRepository) GetByID(id int) (*model.Task, error) {
    var task model.Task
    if err := t.db.Where("id = ?", id).First(&task).Error; err != nil {
        return nil, err
    }
    return &task, nil
}


func (t *taskRepository) Update(id int, task *model.Task) error {
	if err := t.db.Model(&model.Task{}).Where("id = ?", id).Updates(task).Error; err != nil {
		return err
	}
	return nil
}

func (t *taskRepository) Delete(id int) error {
	if err := t.db.Where("id = ?", id).Delete(&model.Task{}).Error; err != nil {
		return err
	}
	return nil
}
