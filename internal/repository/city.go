package repository

import (
	"context"
	"encoding/json"
	"log"
	"time"

	"github.com/OctavianoRyan25/belajar-pattern-code-go/internal/domain"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

type CityRepository interface {
	GetAllCity(context.Context) ([]domain.City, error)
}

type cityRepo struct {
	db    *gorm.DB
	redis *redis.Client
}

func NewCityRepository(db *gorm.DB, redis *redis.Client) CityRepository {
	return &cityRepo{
		db:    db,
		redis: redis,
	}
}

func (r *cityRepo) GetAllCity(ctx context.Context) ([]domain.City, error) {
	var cities []domain.City
	cacheKey := "city:all"
	val, err := r.redis.Get(ctx, cacheKey).Result()
	if err == nil {
		err := json.Unmarshal([]byte(val), &cities)
		if err == nil {
			log.Println("GetAll: Mengambil data dari Redis (Cache Hit)")
			return cities, nil
		}
	}

	if err := r.db.WithContext(ctx).Find(&cities).Error; err != nil {
		return nil, err
	}

	jsonData, err := json.Marshal(cities)
	if err == nil {
		// Set cache dengan TTL (misal 1 jam)
		r.redis.Set(ctx, cacheKey, jsonData, 1*time.Hour)
		log.Println("GetAll: Menyimpan data ke Redis")
	}

	return cities, nil
}
