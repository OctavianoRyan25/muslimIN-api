package domain

import (
	"time"

	"gorm.io/datatypes"
)

type JadwalSholat struct {
	ID        uint `gorm:"primaryKey"`
	City      string
	Schedule  datatypes.JSON
	Date      time.Time
	CreatedAt time.Time
	UpdatedAt time.Time
}
