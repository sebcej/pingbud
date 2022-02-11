package common

import (
	"fmt"
	"os"

	"github.com/spf13/viper"
)

var Settings *viper.Viper

type SettingsAttrs struct {
	PingRoute string `json:"pingroute" binding:"required"`
	PingCron  string `json:"pingcron" binding:"required"`
	Enabled   bool   `json:"enabled" binding:"required"`
	Timeout   int    `json:"timeout" binding:"required"`
	PingCount int    `json:"pingcount" binding:"required"`
	Retention int    `json:"retention" binding:"required"`
}

func setSettingsDefaults() {
	viper.SetDefault("pingroute", SETTINGS_DEFAULTS.PingRoute)
	viper.SetDefault("pingcron", SETTINGS_DEFAULTS.PingCron)
	viper.SetDefault("enabled", SETTINGS_DEFAULTS.Enabled)
	viper.SetDefault("timeout", SETTINGS_DEFAULTS.Timeout)
	viper.SetDefault("pingcount", SETTINGS_DEFAULTS.PingCount)
	viper.SetDefault("retention", SETTINGS_DEFAULTS.Retention)
}

func InitSettings() {
	filePath := os.Getenv("SETTINGS_PATH")
	if filePath == "" {
		filePath = "/data/"
	}

	viper.SetConfigName("pingbud_conf.json")
	viper.SetConfigType("json")
	viper.AddConfigPath(filePath)

	setSettingsDefaults()

	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			fmt.Println("Creating conf")
			viper.WriteConfigAs(filePath + "/" + SETTINGS_FILE_NAME)
		} else {
			panic("Conf error")
		}
	}

	Settings = viper.GetViper()
}

func UpdateSetting(key string, value interface{}) {
	Settings.Set(key, value)

	viper.WriteConfig()
}

func UpdateSettings(settings map[string]interface{}) {
	for key, value := range settings {
		Settings.Set(key, value)
	}

	viper.WriteConfig()
}
