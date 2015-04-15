package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"github.com/Treeptik/docker-viz/flare"
)

func main() {
	r := gin.Default()
	r.LoadHTMLGlob("templates/*")
	r.GET("/", func(c *gin.Context) {
		obj := gin.H{"title": "Main website"}
		c.HTML(http.StatusOK, "dendrogam.tmpl", obj)
	})

	r.GET("/flare.json", func(c *gin.Context) {
		c.String(http.StatusOK, flare.Default())
	})

	// Listen and server on 0.0.0.0:8080
	r.Run(":8080")
}