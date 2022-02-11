package cron

import (
	"pingbud/common"

	"github.com/robfig/cron/v3"
)

var cronInstance *cron.Cron

func Init() {
	cronInstance = cron.New(cron.WithSeconds())

	cronInstance.AddFunc(common.Settings.GetString("pingcron"), pingCron)

	cronInstance.Start()
}

func Restart() {
	cronInstance.Stop()
	Init()
}
