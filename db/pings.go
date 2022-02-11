package db

import (
	"time"

	"github.com/asdine/storm/v3"
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

func InsertPing(data PingTest) error {
	return dbi.From("pings").Save(&data)
}

func GetPings() (pings []PingTest, err error) {
	err = dbi.From("pings").All(&pings, storm.Limit(100))
	return
}

func GetDailyStats() (daily GenericStats, err error) {
	var pings []PingTest
	now := time.Now()
	today := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, &time.Location{}).Unix()

	dbi.From("pings").Select(q.Gte("Time", today), q.Eq("IsOnline", false)).Find(&pings)
	daily.Errors = len(pings)

	dbi.From("pings").Select(q.Gte("Time", today)).Find(&pings)

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
