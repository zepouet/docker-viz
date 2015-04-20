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
	r.GET("/", func(c *gin.Context) {
		obj := gin.H{"title": "Main website"}
		c.HTML(http.StatusOK, "dendrogam.tmpl", obj)
	})

	r.GET("/buble", func(c *gin.Context) {
		obj := gin.H{"title": "Buble"}
		c.HTML(http.StatusOK, "buble.tmpl", obj)
	})

	r.GET("/flare.json", func(c *gin.Context) {
		c.String(http.StatusOK, flare.Dendrogam(&dockerClient))
	})

	// Listen and server on 0.0.0.0:8080
	r.Run(":" + strconv.Itoa(vizPort))
}