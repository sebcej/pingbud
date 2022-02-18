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

	var tempPings []PingTest
	err = dbi.From("pings").Select(q.Gte("Time", start), q.Lte("Time", end)).Find(&tempPings)

	divider := 1

	// Check which divider we need in order to have a reasonable number of elements for the chart
	for {
		if len(tempPings)/divider >= common.MAX_SENDABLE_VALS {
			divider++
			continue
		}
		break
	}

	var groupedPings [][]PingTest
	tempPingsLen := len(tempPings)
	for i := divider; i < tempPingsLen; i += divider {
		groupedPings = append(groupedPings, tempPings[i-divider:i])
	}

	mod := tempPingsLen % divider
	groupedPings = append(groupedPings, tempPings[tempPingsLen-mod:])

	// Aggregate the numbers and get the max and min values of each aggregation
	for _, groupedPing := range groupedPings {
		blockLen := len(groupedPing)
		isOnline := true
		var avg float64
		var min float64
		var max float64
		var jitter float64

		if blockLen == 0 {
			continue
		}

		for index, ping := range groupedPing {
			if !ping.IsOnline {
				isOnline = false
			}
			avg += ping.Avg
			if ping.Min < min || index == 0 {
				min = ping.Min
			}
			if ping.Max > max {
				max = ping.Max
			}
			jitter += ping.Jitter
		}

		pings = append(pings, PingTest{
			Time:     groupedPing[blockLen-1].Time,
			IsOnline: isOnline,
			Avg:      common.ToFixed(avg/float64(blockLen), 6),
			Min:      common.ToFixed(min, 6),
			Max:      common.ToFixed(max, 6),
			Jitter:   common.ToFixed(jitter/float64(blockLen), 6),
		})
	}

	return
}

func GetLatestPings(filter string) (pings []PingTest, err error) {
	start, end := manageFilters(filter)

	dbi.From("pings").Select(q.Gte("Time", start), q.Lte("Time", end)).Limit(20).OrderBy("Time").Reverse().Find(&pings)
	return
}

func GetStats(filter string) (daily GenericStats, err error) {
	var pings []PingTest
	start, end := manageFilters(filter)

	dbi.From("pings").Select(q.Gte("Time", start), q.Lte("Time", end), q.Eq("IsOnline", false)).Find(&pings)
	daily.Errors = len(pings)

	dbi.From("pings").Select(q.Gte("Time", start), q.Lte("Time", end)).Find(&pings)

	pingsLen := len(pings)
	if pingsLen > 0 {
		for _, ping := range pings {
			if !ping.IsOnline {
				continue
			}

			daily.AvgPing += ping.Avg
			daily.AvgJitter += ping.Jitter
		}

		daily.AvgPing = daily.AvgPing / float64(pingsLen)
		daily.AvgJitter = daily.AvgJitter / float64(pingsLen)
	}

	return
}
