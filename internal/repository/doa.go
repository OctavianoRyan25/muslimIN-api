package repository

import (
	"math/rand"

	"github.com/OctavianoRyan25/belajar-pattern-code-go/internal/domain"
	"gorm.io/gorm"
)

type DoaRepository interface {
	GetAll() ([]domain.Doa, error)
	GetById(id uint) (*domain.Doa, error)
	GetRandom() (*domain.Doa, error)
	CountDoa() (uint, error)
}

type doaRepo struct {
	db *gorm.DB
}

func NewDoaRepository(db *gorm.DB) DoaRepository {
	return &doaRepo{db: db}
}

func (r *doaRepo) GetAll() ([]domain.Doa, error) {
	var doas []domain.Doa
	if err := r.db.Find(&doas).Error; err != nil {
		return nil, err
	}
	return doas, nil
}

func (r *doaRepo) GetById(id uint) (*domain.Doa, error) {
	var doa domain.Doa
	if err := r.db.Where("id = ?", id).First(&doa).Error; err != nil {
		return nil, err
	}
	return &doa, nil
}

func (r *doaRepo) GetRandom() (*domain.Doa, error) {
	// Get random by id, based on max
	random, err := r.CountDoa()
	if err != nil {
		return nil, err
	}
	id := rand.Intn(int(random))

	var doa domain.Doa
	if err := r.db.Where("id = ?", id).First(&doa).Error; err != nil {
		return nil, err
	}
	return &doa, nil
}

func (r *doaRepo) CountDoa() (uint, error) {
	var count int64
	if err := r.db.Model(&domain.Doa{}).Count(&count).Error; err != nil {
		return 0, err
	}
	return uint(count), nil
}
