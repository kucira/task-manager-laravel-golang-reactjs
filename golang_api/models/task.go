package models

import (
	"time"
)

type Task struct {
	ID        uint      `json:"id" gorm:"primaryKey:autoIncrement"`
	Name      string    `json:"name" gorm:"type:varchar(100);not null"`
	CreatedAt time.Time `json:"created_at"`
	UpdateAt  time.Time `json:"updated_at"`
}
