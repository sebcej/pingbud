package api

import (
	"net/http"
	"pingbud/db"

	"github.com/gin-gonic/gin"
)

// Chart documentation for annotations generation
// https://apexcharts.com/docs/annotations/

func getStats(c *gin.Context) {
	filter := c.DefaultQuery("filter", "")

	aggregatedPings, err := db.GetAggregatedPings(filter)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	latestPings, err := db.GetLatestPings(filter)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	stats, err := db.GetStats(filter)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	c.JSON(http.StatusOK, gin.H{
		"aggregated": aggregatedPings,
		"latest":     latestPings,
		"stats":      stats,
	})
}
