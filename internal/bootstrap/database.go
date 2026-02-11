package bootstrap

import (
	"fmt"
	"log"

	"github.com/OctavianoRyan25/belajar-pattern-code-go/internal/config"
	"github.com/OctavianoRyan25/belajar-pattern-code-go/internal/domain"
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

	err = db.AutoMigrate(
		&domain.User{},
	)

	if err != nil {
		log.Fatalf("Failed auto migrate: %v", err)
	}

	log.Println("Auto migration completed")

	return db
}
