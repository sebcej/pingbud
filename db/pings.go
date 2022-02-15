package db

import (
	"pingbud/common"
	"time"

	"github.com/asdine/storm/v3/q"
)

type PingTest struct {
	Time     int64           `storm:"id" json:"time"`
	IsOnline bool            `json:"isOnline"`
	Pings    []time.Duration `json:"pings"`
	Avg      float64         `json:"avg"`
	Min      float64         `json:"min"`
	Max      float64         `json:"max"`
	Jitter   float64         `json:"jitter"`
}

type GenericStats struct {
	Errors    int     `json:"errors"`
	AvgJitter float64 `json:"avgJitter"`
	AvgPing   float64 `json:"avgPing"`
}

func manageFilters(filter string) (start int64, end int64) {
	var day time.Time

	if filter == "" {
		day = time.Now()

		end = day.Unix()
		day = day.Add(time.Duration(-24) * time.Hour)
		start = day.Unix()
	} else {
		day, _ = time.Parse("2006-01-02", filter)

		start = time.Date(day.Year(), day.Month(), day.Day(), 0, 0, 0, 0, &time.Location{}).Unix()
		end = time.Date(day.Year(), day.Month(), day.Day(), 23, 59, 59, 9, &time.Location{}).Unix()
	}

	return
}

func InsertPing(data PingTest) error {
	return dbi.From("pings").Save(&data)
}

func GetAggregatedPings(filter string) (pings []PingTest, err error) {
	start, end := manageFilters(filter)

	err = dbi.From("pings").Select(q.Gte("Time", start), q.Lte("Time", end)).Find(&pings)

	divider := 1

	for {
		if len(pings)/divider >= common.MAX_SENDABLE_VALS {
			divider++
			continue
		}
		break
	}

	return
}

func GetLatestPings(filter string) (pings []PingTest, err error) {
	start, end := manageFilters(filter)

	dbi.From("pings").Select(q.Gte("Time", start), q.Lte("Time", end)).Limit(20).OrderBy("Time").Reverse().Find(&pings)
	return
}

func GetDailyStats(filter string) (daily GenericStats, err error) {
	var pings []PingTest
	start, end := manageFilters(filter)

	dbi.From("pings").Select(q.Gte("Time", start), q.Lte("Time", end), q.Eq("IsOnline", false)).Find(&pings)
	daily.Errors = len(pings)

	dbi.From("pings").Select(q.Gte("Time", start), q.Lte("Time", end)).Find(&pings)

	pingsLen := len(pings)
	if pingsLen > 0 {
		for _, ping := range pings {
			daily.AvgPing += ping.Avg
			daily.AvgJitter += ping.Jitter
		}

		daily.AvgPing = daily.AvgPing / float64(pingsLen)
		daily.AvgJitter = daily.AvgJitter / float64(pingsLen)
	}

	return
}
