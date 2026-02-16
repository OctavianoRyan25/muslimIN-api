package seed

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"time"

	"github.com/OctavianoRyan25/belajar-pattern-code-go/internal/delivery/http/request"
	"github.com/OctavianoRyan25/belajar-pattern-code-go/internal/domain"
	"github.com/OctavianoRyan25/belajar-pattern-code-go/internal/public"
	"gorm.io/datatypes"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func SeedData(db *gorm.DB) error {
	basePath := "./data/adzan"

	// Cek apakah folder ada, jika tidak, coba naik satu level (untuk testing lokal)
	if _, err := os.Stat(basePath); os.IsNotExist(err) {
		basePath = "../data/adzan"
	}

	now := time.Now()
	year := now.Format("2006")
	month := now.Format("01")

	// 1. Ambil semua folder kota
	cities, err := os.ReadDir(basePath)
	if err != nil {
		return fmt.Errorf("gagal membaca folder data: %v", err)
	}

	for _, city := range cities {
		if !city.IsDir() {
			continue
		}

		cityName := city.Name()
		filePath := filepath.Join(basePath, cityName, year, month+".json")

		if _, err := os.Stat(filePath); os.IsNotExist(err) {
			fmt.Printf("[Seed] File tidak ditemukan untuk kota %s: %s\n", cityName, filePath)
			continue
		}

		fileByte, err := os.ReadFile(filePath)
		if err != nil {
			fmt.Printf("[Seed] Gagal membaca file %s: %v\n", cityName, err)
			continue
		}

		var adzanRows []request.AdzanRow
		if err := json.Unmarshal(fileByte, &adzanRows); err != nil {
			fmt.Printf("[Seed] Gagal parse JSON %s: %v\n", cityName, err)
			continue
		}

		fmt.Printf("[Seed] Memproses %d data untuk kota: %s\n", len(adzanRows), cityName)

		for _, row := range adzanRows {
			dateParsed, _ := time.Parse("2006-01-02", row.Tanggal)

			scheduleByte, _ := json.Marshal(row)

			jadwal := domain.JadwalSholat{
				City:     cityName,
				Schedule: datatypes.JSON(scheduleByte),
				Date:     dateParsed,
			}

			db.Clauses(clause.OnConflict{
				Columns:   []clause.Column{{Name: "city"}, {Name: "date"}},
				DoUpdates: clause.AssignmentColumns([]string{"schedule", "updated_at"}),
			}).Create(&jadwal)
		}
	}

	return nil
}

func SeedCities(db *gorm.DB) error {
	var count int64
	db.Model(&domain.City{}).Count(&count)

	if count == 0 {
		var cities []domain.City
		json.Unmarshal(public.CityJSON, &cities)

		db.CreateInBatches(cities, 100)
		log.Println("Cities seeded successfully!")
	}

	return nil
}
