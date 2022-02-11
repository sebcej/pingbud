package cron

import (
	"pingbud/common"
	"pingbud/db"
	"time"

	"github.com/go-ping/ping"
)

func pingCron() {
	if !common.Settings.GetBool("enabled") {
		return
	}

	error := false

	pinger, err := ping.NewPinger(common.Settings.GetString("pingroute"))
	if err != nil {
		panic(err)
	}
	pinger.Timeout = time.Duration(common.Settings.GetInt("timeout")) * time.Second
	pinger.Count = common.Settings.GetInt("pingcount")
	err = pinger.Run()
	if err != nil {
		error = true
	}
	stats := pinger.Statistics()

	parsedStats := db.PingTest{
		Time:     time.Now().Unix(),
		IsOnline: !error,
		Pings:    stats.Rtts,
		Max:      float64(stats.MaxRtt.Nanoseconds()) / 1000000,
		Min:      float64(stats.MinRtt.Nanoseconds()) / 1000000,
		Avg:      float64(stats.AvgRtt.Nanoseconds()) / 1000000,
		Jitter:   float64(stats.StdDevRtt.Nanoseconds()) / 1000000,
	}

	db.InsertPing(parsedStats)
}
