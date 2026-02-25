package usecase

import (
	"context"

	"github.com/OctavianoRyan25/belajar-pattern-code-go/internal/domain"
	"github.com/OctavianoRyan25/belajar-pattern-code-go/internal/repository"
)

type CityUseCase interface {
	GetAll(context.Context) ([]domain.City, error)
}

type cityUseCase struct {
	repo repository.CityRepository
}

func NewCityUsecase(repo repository.CityRepository) CityUseCase {
	return &cityUseCase{repo: repo}
}

func (u *cityUseCase) GetAll(ctx context.Context) ([]domain.City, error) {
	return u.repo.GetAllCity(ctx)
}
