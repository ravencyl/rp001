package models

import (
	"time"
)

type Task struct {
	ID          uint   `gorm:"primaryKey,index"`
	Name        string `json:"name"`
	Content     string `gorm:"text" json:"content"`
	Status      string
	TargetAt    *time.Time
	CompletedAt *time.Time
	CreatedAt   time.Time `gorm:"autoCreateTime"`
	UpdatedAt   time.Time `gorm:"autoCreateTime"`
}
