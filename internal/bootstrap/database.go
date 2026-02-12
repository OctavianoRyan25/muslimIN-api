package bootstrap

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"strings"

	"github.com/OctavianoRyan25/belajar-pattern-code-go/internal/config"
	"github.com/OctavianoRyan25/belajar-pattern-code-go/internal/delivery/http/request"
	"github.com/OctavianoRyan25/belajar-pattern-code-go/internal/domain"
	"github.com/OctavianoRyan25/belajar-pattern-code-go/internal/public"
	"github.com/OctavianoRyan25/belajar-pattern-code-go/internal/repository"
	"github.com/redis/go-redis/v9"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func InitDatabase(cfg *config.Config) *gorm.DB {
	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		cfg.DBUser,
		cfg.DBPassword,
		cfg.DBHost,
		cfg.DBPort,
		cfg.DBName,
	)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})

	if err != nil {
		log.Fatalf("Failed to connect database: %v", err)
	}

	log.Println("Database connected")

	// AutoMigrate
	err = db.AutoMigrate(
		&domain.User{},
		&domain.APIKey{},
		&domain.Doa{},
	)

	if err != nil {
		log.Fatalf("Failed auto migrate: %v", err)
	}

	log.Println("Auto migration completed")

	return db
}

func seedDoa(db *gorm.DB, redis *redis.Client, ctx context.Context) error {
	repo := repository.NewDoaRepository(db, redis)
	count, err := repo.CountDoa(ctx)

	if err != nil {
		return fmt.Errorf("gagal menghitung data doa: %w", err)
	}

	if count > 0 {
		log.Print("Nothing to Seed")
		return nil
	}

	// processed Seeding
	log.Print("Seeding data dari JSON...")

	var doaReqs []request.Doa
	err = json.Unmarshal(public.DoaJSON, &doaReqs)
	if err != nil {
		return fmt.Errorf("gagal unmarshal JSON: %w", err)
	}

	var doas []domain.Doa
	for _, v := range doaReqs {
		doas = append(doas, domain.Doa{
			Nama:          v.Nama,
			Lafal:         v.Lafal,
			Transliterasi: v.Transliterasi,
			Arti:          v.Arti,
			Riwayat:       v.Riwayat,
			Keterangan:    &v.Keterangan,
			KataKunci:     strings.Join(v.KataKunci, ","),
		})
	}
	if err := db.Create(&doas).Error; err != nil {
		return fmt.Errorf("gagal menyimpan data seed ke database: %w", err)
	}

	log.Printf("Berhasil seeding %d data doa.", len(doas))
	return nil
}
