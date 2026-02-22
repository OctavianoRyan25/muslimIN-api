package repository

import (
	"time"

	"github.com/OctavianoRyan25/belajar-pattern-code-go/internal/domain"
	"gorm.io/gorm"
)

type JadwalSholatRepository interface {
	GetToday(city string) (*domain.JadwalSholat, error)
	GetByDate(city string, date time.Time) (*domain.JadwalSholat, error)
}

type jadwalSholatRepo struct {
	db *gorm.DB
}

func NewJadwalSholatRepository(db *gorm.DB) JadwalSholatRepository {
	return &jadwalSholatRepo{db: db}
}

func (r *jadwalSholatRepo) GetToday(city string) (*domain.JadwalSholat, error) {
	var jadwalSholat domain.JadwalSholat
	now := time.Now()
	startOfDay := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, now.Location())
	endOfDay := startOfDay.Add(24 * time.Hour).Add(-time.Second)

	err := r.db.Where("city = ?", city).Where("date BETWEEN ? AND ?", startOfDay, endOfDay).First(&jadwalSholat).Error

	if err != nil {
		return nil, err
	}

	return &jadwalSholat, nil
}

func (r *jadwalSholatRepo) GetByDate(city string, date time.Time) (*domain.JadwalSholat, error) {
	var jadwalSholat domain.JadwalSholat
	startOfDay := time.Date(date.Year(), date.Month(), date.Day(), 0, 0, 0, 0, date.Location())
	endOfDay := startOfDay.Add(24 * time.Hour).Add(-time.Second)

	err := r.db.Where("city = ?", city).Where("date BETWEEN ? AND ?", startOfDay, endOfDay).First(&jadwalSholat).Error

	if err != nil {
		return nil, err
	}

	return &jadwalSholat, nil
}
