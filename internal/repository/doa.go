package repository

import (
	"context"
	"log"
	"math/rand"
	"strconv"
	"time"

	"github.com/OctavianoRyan25/belajar-pattern-code-go/internal/domain"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

type DoaRepository interface {
	GetAll() ([]domain.Doa, error)
	GetById(id uint) (*domain.Doa, error)
	GetRandom(context.Context) (*domain.Doa, error)
	CountDoa(context.Context) (uint, error)
}

type doaRepo struct {
	db    *gorm.DB
	redis *redis.Client
}

func NewDoaRepository(db *gorm.DB, redis *redis.Client) DoaRepository {
	return &doaRepo{
		db:    db,
		redis: redis,
	}
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

func (r *doaRepo) GetRandom(ctx context.Context) (*domain.Doa, error) {
	// Get random by id, based on max
	random, err := r.CountDoa(ctx)
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

func (r *doaRepo) CountDoa(ctx context.Context) (uint, error) {
	val, err := r.redis.Get(ctx, "doa:count").Result()
	if err == nil {
		count, _ := strconv.Atoi(val)
		return uint(count), nil
	}

	var count int64
	if err := r.db.WithContext(ctx).Model(&domain.Doa{}).Count(&count).Error; err != nil {
		return 0, err
	}

	err = r.redis.Set(ctx, "doa:count", count, 24*time.Hour).Err()
	if err != nil {
		log.Printf("Gagal menyimpan cache ke Redis: %v", err)
	}

	return uint(count), nil
}
