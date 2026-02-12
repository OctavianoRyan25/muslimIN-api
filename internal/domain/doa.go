package domain

import "time"

type Doa struct {
	ID            uint `gorm:"primaryKey"`
	Nama          string
	Lafal         string
	Transliterasi string
	Arti          string
	Riwayat       string
	Keterangan    *string
	KataKunci     string
	CreatedAt     time.Time
	UpdatedAt     time.Time
}
