package main

import (
	"embed"
	"pingbud/api"
	"pingbud/common"
	"pingbud/cron"
	"pingbud/db"

	_ "github.com/joho/godotenv/autoload"
)

//go:embed frontend/dist/spa/*
var static embed.FS

func main() {
	db.Init()
	common.InitSettings()
	cron.Init()
	api.Init(static)
}
