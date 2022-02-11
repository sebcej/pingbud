package common

const (
	SETTINGS_FILE_NAME = "pingbud_conf.json"
)

var (
	SETTINGS_DEFAULTS = SettingsAttrs{
		PingRoute: "8.8.8.8",
		PingCron:  "0 * * * * *",
		Enabled:   false,
		Timeout:   30,
		PingCount: 3,
		Retention: 30,
	}
)
