package models

import "time"

type Task struct {
	ID uint `json:"id" gorm:"primaryKey"`
	Title string `json:"title"`
	Description string `json:"description"`
	IsCompleted bool `json:"isCompleted"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

type CreateTask struct {
	Title string `json:"title" binding:"required"`
	Description string `json:"description"`
	IsCompleted string `json:"isCompleted"`
}