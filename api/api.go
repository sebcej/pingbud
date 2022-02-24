package api

import (
	"embed"
	"io/fs"
	"net/http"
	"os"

	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
)

// https://github.com/gin-gonic/gin/issues/75
// https://github.com/gin-contrib/static/issues/19#issuecomment-800719193

type embedFileSystem struct {
	http.FileSystem
}

func (e embedFileSystem) Exists(prefix string, path string) bool {
	_, err := e.Open(path)
	return err == nil
}

func stripNEmbed(root fs.FS) static.ServeFileSystem {
	stripped, _ := fs.Sub(root, "frontend/dist/spa")

	return embedFileSystem{
		FileSystem: http.FS(stripped),
	}
}

func Init(staticRoot embed.FS) {
	address := os.Getenv("PINGBUD_WEB_ADDRESS")
	if address == "" {
		address = ":8080"
	}

	r := gin.Default()

	r.Use(static.Serve("/", stripNEmbed(staticRoot)))

	v1Api := r.Group("/api/v1")
	{
		v1Api.GET("/settings", getSettings)
		v1Api.POST("/settings", setSettings)
		v1Api.GET("/stats", getStats)
	}

	r.Run(address)
}
