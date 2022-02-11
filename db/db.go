package db

import (
	"os"

	"github.com/asdine/storm/v3"
)

var dbi *storm.DB

func Init() *storm.DB {
	dbPath := os.Getenv("DB_PATH")

	db, err := storm.Open(dbPath + "/pingbud.db")
	if err != nil {
		panic(err)
	}

	dbi = db

	return db
}
