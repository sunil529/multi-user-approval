package models

import "time"

type Task struct {
	ID          uint      `gorm:"primaryKey"`
	Title       string    `gorm:"not null"`
	Description string
	Status      string    `gorm:"default:Pending"`
	CreatedBy   uint      `gorm:"not null"`
	CreatedAt   time.Time
}
