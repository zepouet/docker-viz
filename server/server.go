package server

import (
	"net/http"
	"strconv"
	"html/template"
	"github.com/gin-gonic/gin"
	"github.com/Treeptik/docker-viz/flare"
	"github.com/Treeptik/docker-viz/dockertype"
)

// Function who start the web server of Docker-Viz
func StartServer(port int) {
	r := gin.Default()

	// create static route for all files in this folder
	r.Static("/images", "./asset/images")
	r.Static("/js", "./asset/js")
	r.Static("/fonts", "./asset/fonts")
	r.Static("/css", "./asset/css")
	var baseTemplate = "templates/"

	// index page
	r.GET("/", func(c *gin.Context) {
		countImg := dockertype.CountDockerImages()
		countCont := dockertype.CountDockerContainer()
		obj := gin.H{"title": "Dashboard", "countImages": strconv.Itoa(countImg), "countContainers": strconv.Itoa(countCont)}
		r.SetHTMLTemplate(template.Must(template.ParseFiles(baseTemplate + "main.tpl", baseTemplate + "index.tpl")))
		c.HTML(http.StatusOK, "base", obj)
	})

	// docker success connexion
	r.GET("/docker", func(c *gin.Context) {
		if dockertype.DockerStatut() {
			c.String(http.StatusOK, "true")
		} else {
			c.String(http.StatusOK, "false")
		}
	})

	// dendrogam diagram page
	r.GET("/dendrogam", func(c *gin.Context) {
		obj := gin.H{"title": "Dendrogam Images", "type": "images"}
		r.SetHTMLTemplate(template.Must(template.ParseFiles(baseTemplate + "main.tpl", baseTemplate + "dendrogam.tpl")))
		c.HTML(http.StatusOK, "base", obj)
	})

	// bubble diagram page
	r.GET("/bubble/:name", func(c *gin.Context) {
		var obj gin.H
		switch name := c.Params.ByName("name"); name {
		case "images":
			obj = gin.H{"title": "Buble Container", "type": name}
		case "containers":
			obj = gin.H{"title": "Buble Container", "type": name}
		default:
			c.String(http.StatusNotFound, "404 page not found")
		}
		r.SetHTMLTemplate(template.Must(template.ParseFiles(baseTemplate + "main.tpl", baseTemplate + "bubble.tpl")))
		c.HTML(http.StatusOK, "base", obj)
	})

	// dendrogam diagram page
	r.GET("/miserables", func(c *gin.Context) {
		obj := gin.H{"title": "Miserables Images", "type": "containers"}
		r.SetHTMLTemplate(template.Must(template.ParseFiles(baseTemplate + "main.tpl", baseTemplate + "miserables.tpl")))
		c.HTML(http.StatusOK, "base", obj)
	})

	// json for all diagram route
	r.GET("/json/:name", func(c *gin.Context) {
		switch name := c.Params.ByName("name"); name {
		case "dendrogam":
			c.String(http.StatusOK, flare.DendrogamFlare())
		case "miserables":
			c.String(http.StatusOK, flare.MiserablesFlare())
		default:
			c.String(http.StatusNotFound, "404 page not found")
		}
	})

	// json for bubble diagram special route
	r.GET("/json/:name/:who", func(c *gin.Context) {
		name := c.Params.ByName("name")
		switch who := c.Params.ByName("who"); who {
		case "bubble":
			c.String(http.StatusOK, flare.BubbleFlare(name))
		default:
			c.String(http.StatusNotFound, "404 page not found")
		}
	})

	// start server (default port 8080)
	r.Run(":" + strconv.Itoa(port))
}