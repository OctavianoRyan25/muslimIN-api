package domain

import "time"

type APIKey struct {
	ID           uint `gorm:"primaryKey"`
	UserID       uint
	Key          string `gorm:"uniqueIndex"`
	MonthlyLimit int
	UsageCount   int
	ResetAt      time.Time
	CreatedAt    time.Time
}
