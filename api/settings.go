package api

import (
	"net/http"
	"pingbud/common"
	"pingbud/cron"

	"github.com/gin-gonic/gin"
)

func getSettings(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"settings": common.Settings.AllSettings(),
	})
}

func setSettings(c *gin.Context) {
	// Validate the settings
	var settings common.SettingsAttrs
	if err := c.ShouldBindJSON(&settings); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	common.UpdateSettings(settings, false)

	cron.Restart()

	c.JSON(http.StatusOK, gin.H{
		"success": true,
	})
}
