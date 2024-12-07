package models

import "time"

type TaskApproval struct {
	ID         uint      `gorm:"primaryKey"`
	TaskID     uint      `gorm:"not null"`
	ApprovedBy uint      `gorm:"not null"`
	Comment    string
	ApprovedAt time.Time
}
