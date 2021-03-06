package router

import (
	"github.com/TonyDoen/go_code_review/gin-toolkit/service/demo"
	"github.com/robfig/cron"
)

func initCron() (c *cron.Cron) {
	c = cron.New()
	c.Start()
	return c
}

func Crontab() {
	c := initCron()

	_ = c.AddFunc("@every 1s", demo.CronCall)
	_ = c.AddFunc("*/1 * * * * *", demo.CronCall)
}
