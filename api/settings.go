package api

import (
	"net/http"
	"pingbud/common"

	"github.com/gin-gonic/gin"
)

func getSettings(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"settings": common.Settings.AllSettings(),
	})
}

func setSettings(c *gin.Context) {
	// Validate the settings
	var settingsValidator common.SettingsAttrs
	if err := c.ShouldBindJSON(&settingsValidator); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Sent them in a way that supports Viper
	var settings map[string]interface{}
	c.BindJSON(&settings)
	common.UpdateSettings(settings)
}
