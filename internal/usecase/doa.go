package usecase

import (
	"context"
	"errors"

	"github.com/OctavianoRyan25/belajar-pattern-code-go/internal/domain"
	"github.com/OctavianoRyan25/belajar-pattern-code-go/internal/repository"
)

type DoaUseCase interface {
	GetAll() ([]domain.Doa, error)
	GetById(id uint) (*domain.Doa, error)
	GetRandom(ctx context.Context) (*domain.Doa, error)
}

type doaUseCase struct {
	repo repository.DoaRepository
}

func NewDoaUsecase(repo repository.DoaRepository) DoaUseCase {
	return &doaUseCase{repo: repo}
}

func (u *doaUseCase) GetAll() ([]domain.Doa, error) {
	return u.repo.GetAll()
}
func (u *doaUseCase) GetById(id uint) (*domain.Doa, error) {
	if id == 0 {
		return nil, errors.New("ID doa tidak valid")
	}
	return u.repo.GetById(id)
}
func (u *doaUseCase) GetRandom(ctx context.Context) (*domain.Doa, error) {
	return u.repo.GetRandom(ctx)
}
