package cron

import (
	"github.com/robfig/cron/v3"
	"gorm.io/gorm"
)

type CronJob struct {
	Scheduler *cron.Cron
	DB        *gorm.DB
}

func NewCronJob(db *gorm.DB) *CronJob {
	return &CronJob{
		Scheduler: cron.New(),
		DB:        db,
	}
}

func (c *CronJob) Start() {
	c.RegisterTasks()
	c.Scheduler.Start()
}
