package main

import (
	"github.com/gin-gonic/gin"
	"github.com/tomazk/envcfg"
	"net/http"
	"docker-viz/flare"
	"strconv"
)

type Config struct {
	VIZ_PORT int
	DOCKER_CLIENT string
}

func LoadConfig() (int, string){
	var cfg Config
	envcfg.Unmarshal(&cfg)
	vizPort :=  cfg.VIZ_PORT
	dockerClient := cfg.DOCKER_CLIENT

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

	r.GET("/flare.json", func(c *gin.Context) {
		c.String(http.StatusOK, flare.Dendrogam(dockerClient))
	})

	// Listen and server on 0.0.0.0:8080
	r.Run(":" + strconv.Itoa(vizPort))
}