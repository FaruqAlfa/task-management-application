package model

import (
	"gorm.io/gorm"
)

type Task struct {
	gorm.Model
	Title     string  `json:"title"`
	Content   string  `json:"content"`
	UserID    uint    `json:"user_id"`
	User      User    `json:"user"`
	ProjectID uint    `json:"project_id"`
	Project   Project `json:"project"`
}
