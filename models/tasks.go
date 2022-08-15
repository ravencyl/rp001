package models

import "time"

type Task struct {
	ID      uint   `gorm:"primaryKey,index"`
	Name    string `json:"project_name"`
	Summary string `json:"project_summary"`
	Created time.Time
	Updated time.Time
}
