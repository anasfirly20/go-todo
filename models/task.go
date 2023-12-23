package models

import "time"

type Task struct {
	ID uint `json:"id" gorm:"primaryKey"`
	Title string `json:"title"`
	IsCompleted bool `json:"isCompleted"`
	Description string `json:"description"`
	CreatedAt time.Time
	UpdatedAt time.Time
}