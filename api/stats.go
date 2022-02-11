package api

import (
	"net/http"
	"pingbud/db"

	"github.com/gin-gonic/gin"
)

func getStats(c *gin.Context) {
	pings, err := db.GetPings()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	stats, err := db.GetDailyStats()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	c.JSON(http.StatusOK, gin.H{
		"results": pings,
		"stats":   stats,
	})
}
