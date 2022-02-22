package db

import (
	"os"
	"pingbud/common"
	"time"

	"github.com/asdine/storm/v3"
	"github.com/asdine/storm/v3/q"
	"github.com/robfig/cron/v3"
)

var dbi *storm.DB

func Init() *storm.DB {
	dbPath := os.Getenv("DB_PATH")

	db, err := storm.Open(dbPath + "/pingbud.db")
	if err != nil {
		panic(err)
	}

	dbi = db

	cleanupInstance := cron.New()
	cleanupInstance.AddFunc("0 0 * * *", cleanup)

	return db
}

// Perform cleanup of pings database bucket
func cleanup() {
	days := common.Settings.GetInt("retention")
	day := time.Now()
	day.AddDate(0, 0, -days)

	timestamp := time.Date(day.Year(), day.Month(), day.Day(), 0, 0, 0, 0, &time.Location{}).Unix()

	dbi.Select(q.Lte("Time", timestamp)).Delete(new(PingTest))
}
