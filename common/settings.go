package common

import (
	"fmt"
	"os"
	"reflect"

	"github.com/spf13/viper"
)

var Settings *viper.Viper

type SettingsAttrs struct {
	PingRoute      string `json:"pingroute" binding:"required"`
	PingCron       string `json:"pingcron" binding:"required"`
	Enabled        bool   `json:"enabled" binding:""`
	PrivilegedMode bool   `json:"privilegedmode" binding:""`
	Timeout        int    `json:"timeout" binding:"required"`
	PingCount      int    `json:"pingcount" binding:"required"`
	Retention      int    `json:"retention" binding:"required"`
}

func InitSettings() {
	filePath := os.Getenv("SETTINGS_PATH")
	if filePath == "" {
		filePath = "/data/"
	}

	viper.SetConfigName("pingbud_conf.json")
	viper.SetConfigType("json")
	viper.AddConfigPath(filePath)

	Settings = viper.GetViper()

	UpdateSettings(SETTINGS_DEFAULTS, true)

	if err := Settings.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			fmt.Println("Creating conf")
			viper.WriteConfigAs(filePath + "/" + SETTINGS_FILE_NAME)
		} else {
			panic("Conf error")
		}
	}
}

func UpdateSetting(key string, value interface{}) {
	Settings.Set(key, value)

	viper.WriteConfig()
}

// Update settings based on the content of the SettingsAttrs config
func UpdateSettings(settings SettingsAttrs, isDefault bool) {
	attrs := reflect.ValueOf(settings)
	attrType := reflect.TypeOf(settings)

	for i := 0; i < attrs.NumField(); i++ {
		field := attrType.Field(i)

		fieldName := field.Tag.Get("json")
		fieldValue := attrs.Field(i).Interface()

		if isDefault {
			Settings.SetDefault(fieldName, fieldValue)
		} else {
			Settings.Set(fieldName, fieldValue)
		}
	}
	if !isDefault {
		Settings.WriteConfig()
	}
}
