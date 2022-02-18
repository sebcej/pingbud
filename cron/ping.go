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

	max := float64(stats.MaxRtt.Nanoseconds())
	min := float64(stats.MinRtt.Nanoseconds())
	avg := float64(stats.AvgRtt.Nanoseconds())
	jitter := float64(stats.StdDevRtt.Nanoseconds())

	// If data is broken or without any sens we will flag them as an error connection
	// This will exclude them from stats
	if min <= 0 || max <= 0 || avg <= 0 || jitter < 0 {
		error = true

		min = 0
		max = 0
		avg = 0
		jitter = 0
	}

	parsedStats := db.PingTest{
		Time:     time.Now().Unix(),
		IsOnline: !error,
		Pings:    stats.Rtts,
		Max:      max / 1000000,
		Min:      min / 1000000,
		Avg:      avg / 1000000,
		Jitter:   jitter / 1000000,
	}

	db.InsertPing(parsedStats)
}
