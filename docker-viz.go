package main

import (
	"github.com/gin-gonic/gin"
	"github.com/tomazk/envcfg"
	"net/http"
	"strconv"
	"github.com/Treeptik/docker-viz/flare"
	"html/template"
)

type (
	// Represents the system environment variable used for config docker-viz
	Config struct {
		VIZ_PORT int
		DOCKER_HOST string
	}
)

// Load the system environment variable
func LoadConfig() (int, string){
	var cfg Config
	envcfg.Unmarshal(&cfg)
	vizPort :=  cfg.VIZ_PORT
	dockerClient := cfg.DOCKER_HOST

	if vizPort == 0 {
		vizPort = 8080
	}
	if dockerClient == "" {
		dockerClient = "unix:///var/run/docker.sock"
	}

	return vizPort, dockerClient
}

func main() {
	vizPort, dockerClient := LoadConfig()

	r := gin.Default()
	r.Static("/images", "./asset/images")
	r.Static("/js", "./asset/js")
	r.Static("/css", "./asset/css")
	var baseTemplate = "templates/"

	// index page
	r.GET("/", func(c *gin.Context) {
		obj := gin.H{"title": "Index"}
		r.SetHTMLTemplate(template.Must(template.ParseFiles(baseTemplate + "main.tpl", baseTemplate + "index.tpl")))
		c.HTML(http.StatusOK, "base", obj)
	})

	r.GET("/dendrogam", func(c *gin.Context) {
		obj := gin.H{"title": "Dendrogam Images", "type": "images"}
		r.SetHTMLTemplate(template.Must(template.ParseFiles(baseTemplate + "main.tpl", baseTemplate + "dendrogam.tpl")))
		c.HTML(http.StatusOK, "base", obj)
	})

	r.GET("/bubble/:name", func(c *gin.Context) {
		var obj gin.H
		switch name := c.Params.ByName("name"); name {
			case "images":
				obj = gin.H{"title": "Buble Container", "type": "images"}
			case "containers":
				obj = gin.H{"title": "Buble Container", "type": "containers"}
			default:
			c.String(http.StatusNotFound, "404 page not found")
		}
		r.SetHTMLTemplate(template.Must(template.ParseFiles(baseTemplate + "main.tpl", baseTemplate + "bubble.tpl")))
		c.HTML(http.StatusOK, "base", obj)
	})

	r.GET("/flare/:name/json", func(c *gin.Context) {
		switch name := c.Params.ByName("name"); name {
			case "images":
				c.String(http.StatusOK, flare.DendrogamAndBubbleImages(&dockerClient))
			case "containers":
				c.String(http.StatusOK, flare.BubbleContainers(&dockerClient))
			default:
				c.String(http.StatusNotFound, "404 page not found")
		}
	})

	// start server
	r.Run(":" + strconv.Itoa(vizPort))
}