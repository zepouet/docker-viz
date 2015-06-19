package server

import (
	"strconv"
	"github.com/gin-gonic/gin"
)

const(
	TEMPLATES = "asset/templates/"
	TEMPLATE_MAIN    = TEMPLATES + "main.tpl"
)

// Function who start the web server of Docker-Viz
func StartServer(port int, debug bool) {

	if (debug) {
		gin.SetMode(gin.DebugMode)
	} else {
		gin.SetMode(gin.ReleaseMode)
	}

	r := gin.Default()

	// create static route for all files in this folder
	r.Static("/images", "./asset/images")
	r.Static("/js", "./asset/js")
	r.Static("/fonts", "./asset/fonts")
	r.Static("/css", "./asset/css")

	// add Index route
	rIndexController := IndexController{}
	rIndexController.Run(r)

	// add Bubble route
	rBubbleController := BubbleController{}
	rBubbleController.Run(r)

	// add Dendrogam route
	rDendrogramController := DendrogramController{}
	rDendrogramController.Run(r)

	// add Miserables route
	rMiserablesController := MiserablesController{}
	rMiserablesController.Run(r)

	// add Json route
	rJsonController := JsonController{}
	rJsonController.Run(r)

	// add Docker verfication route
	rDockerController := DockerController{}
	rDockerController.Run(r)

	// start server (default port 8080)
	r.Run(":" + strconv.Itoa(port))
}