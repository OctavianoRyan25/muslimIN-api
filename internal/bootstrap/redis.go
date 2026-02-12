package bootstrap

import (
	"context"
	"time"

	"github.com/OctavianoRyan25/belajar-pattern-code-go/internal/config"
	"github.com/redis/go-redis/v9"
)

func InitRedis(cfg *config.Redis) (*redis.Client, error) {
	rdb := redis.NewClient(&redis.Options{
		Addr:     cfg.Addr,
		Password: cfg.Password,
		DB:       int(cfg.DB),
	})

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := rdb.Ping(ctx).Err(); err != nil {
		return nil, err
	}

	return rdb, nil
}
