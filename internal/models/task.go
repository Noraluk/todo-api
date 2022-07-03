package models

import (
	"time"
	"todo-api/internal/common"
)

type Task struct {
	Id          int       `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"created_at" gorm:"<-:create"`
	UpdatedAt   time.Time `json:"updated_at"`
}

func (Task) TableName() string {
	return "todos"
}

type TaskSearchModel struct {
	common.SearchModel
}

type TaskRequest struct {
	Name string `json:"name"`
}
