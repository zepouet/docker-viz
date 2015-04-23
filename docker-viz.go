package main

import (
	"github.com/gin-gonic/gin"
	"github.com/tomazk/envcfg"
	"net/http"
	"strconv"
	"github.com/Treeptik/docker-viz/flare"
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
	r.LoadHTMLGlob("templates/*")

	// index page
	r.GET("/", func(c *gin.Context) {
		obj := gin.H{"title": "Main website"}
		c.HTML(http.StatusOK, "index.tpl", obj)
	})

	r.GET("/dendrogam", func(c *gin.Context) {
		obj := gin.H{"title": "Dendrogam Images", "type": "images"}
		c.HTML(http.StatusOK, "dendrogam.tpl", obj)
	})

	r.GET("/bubble/:name", func(c *gin.Context) {
		switch name := c.Params.ByName("name"); name {
			case "images":
				obj := gin.H{"title": "Buble Container", "type": name}
				c.HTML(http.StatusOK, "bubble.tpl", obj)
			case "containers":
				obj := gin.H{"title": "Buble Container", "type": name}
				c.HTML(http.StatusOK, "bubble.tpl", obj)
			default:
			c.String(http.StatusInternalServerError, "500 Internal Server Error")
		}
	})

	r.GET("/flare/:name/json", func(c *gin.Context) {
		switch name := c.Params.ByName("name"); name {
			case "images":
				c.String(http.StatusOK, flare.DendrogamAndBubbleImages(&dockerClient))
			case "containers":
				c.String(http.StatusOK, flare.BubbleContainers(&dockerClient))
			default:
				c.String(http.StatusInternalServerError, "500 Internal Server Error")
		}
	})

	// start server
	r.Run(":" + strconv.Itoa(vizPort))
}