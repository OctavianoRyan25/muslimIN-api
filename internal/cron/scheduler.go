package cron

import (
	"fmt"

	"github.com/OctavianoRyan25/belajar-pattern-code-go/internal/tasks"
)

func (c *CronJob) RegisterTasks() {
	// 1 monthly sync task
	_, err := c.Scheduler.AddFunc("@monthly", func() {
		fmt.Println("[Cron] Memulai sinkronisasi data bulanan...")

		tasks.ExecuteMonthlySync(c.DB)
	})

	if err != nil {
		fmt.Printf("[Cron] Gagal mendaftarkan task bulanan: %v\n", err)
	}
}
