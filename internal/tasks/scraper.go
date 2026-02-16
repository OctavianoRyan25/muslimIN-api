package tasks

import (
	"log"
	"os"
	"os/exec"

	"github.com/OctavianoRyan25/belajar-pattern-code-go/internal/seed"
	"gorm.io/gorm"
)

func ExecuteMonthlySync(db *gorm.DB) {
	cmd := exec.Command("python", "../python/parser.py")

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		log.Println("Seeding gagal:", err)
	}

	if err := seed.SeedData(db); err != nil {
		log.Println("Seeding gagal:", err)
	}

}
