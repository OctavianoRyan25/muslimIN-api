package domain

import "time"

type APIKey struct {
	ID           uint `gorm:"primaryKey"`
	UserID       uint
	Key          string `gorm:"type:varchar(64);uniqueIndex"`
	MonthlyLimit int
	UsageCount   int
	ResetAt      time.Time
	CreatedAt    time.Time
}
