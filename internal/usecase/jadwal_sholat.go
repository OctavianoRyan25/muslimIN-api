package usecase

import (
	"github.com/OctavianoRyan25/belajar-pattern-code-go/internal/domain"
	"github.com/OctavianoRyan25/belajar-pattern-code-go/internal/repository"
)

type JadwalSholatUseCase interface {
	GetToday() (*domain.JadwalSholat, error)
}

type jadwalSholatUseCase struct {
	repo repository.JadwalSholatRepository
}

func NewJadwalSholatuseCase(repo repository.JadwalSholatRepository) JadwalSholatUseCase {
	return &jadwalSholatUseCase{repo: repo}
}

func (u *jadwalSholatUseCase) GetToday() (*domain.JadwalSholat, error) {
	return u.repo.GetToday()
}
