package usecase

import (
	"time"

	"github.com/OctavianoRyan25/belajar-pattern-code-go/internal/domain"
	"github.com/OctavianoRyan25/belajar-pattern-code-go/internal/repository"
)

type JadwalSholatUseCase interface {
	GetToday(city string) (*domain.JadwalSholat, error)
	GetByDate(city string, date time.Time) (*domain.JadwalSholat, error)
}

type jadwalSholatUseCase struct {
	repo repository.JadwalSholatRepository
}

func NewJadwalSholatuseCase(repo repository.JadwalSholatRepository) JadwalSholatUseCase {
	return &jadwalSholatUseCase{repo: repo}
}

func (u *jadwalSholatUseCase) GetToday(city string) (*domain.JadwalSholat, error) {
	return u.repo.GetToday(city)
}

func (u *jadwalSholatUseCase) GetByDate(city string, date time.Time) (*domain.JadwalSholat, error) {
	return u.repo.GetByDate(city, date)
}
