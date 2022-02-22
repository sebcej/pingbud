package cron

import (
	"pingbud/common"
	"pingbud/db"
	"time"

	"github.com/go-ping/ping"
)

func parseNS(stat time.Duration) float64 {
	return float64(stat.Nanoseconds()) / 1000000
}

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

	max := parseNS(stats.MaxRtt)
	min := parseNS(stats.MinRtt)
	avg := parseNS(stats.AvgRtt)
	jitter := parseNS(stats.StdDevRtt)

	// If data is broken or without any sense we will flag them as an error connection
	// This will exclude them from stats
	if min <= 0 || max <= 0 || avg <= 0 || jitter < 0 {
		error = true

		min = 0
		max = 0
		avg = 0
		jitter = 0
	}

	var pings []float64
	for _, rtt := range stats.Rtts {
		pings = append(pings, parseNS(rtt))
	}

	parsedStats := db.PingTest{
		Time:     time.Now().Unix(),
		IsOnline: !error,
		Pings:    pings,
		Max:      max,
		Min:      min,
		Avg:      avg,
		Jitter:   jitter,
	}

	db.InsertPing(parsedStats)
}
