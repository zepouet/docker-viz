package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()
	r.LoadHTMLGlob("templates/*")
	r.GET("/", func(c *gin.Context) {
		obj := gin.H{"title": "Main website"}
		c.HTML(http.StatusOK, "index_dendrogam.tmpl", obj)
	})

	r.GET("/flare.json", func(c *gin.Context) {
		c.String(http.StatusOK, "json")
	})

	// Listen and server on 0.0.0.0:8080
	r.Run(":8080")
}