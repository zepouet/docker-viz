package server

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"github.com/Treeptik/docker-viz/flare"
)

type JsonController struct {}

/* This method inits all the routes and deletegates to the associated resources component the gathering of data */
func (b *JsonController) Run(router *gin.Engine) {

	// json for all diagram route
	router.GET("/json/:name", func(c *gin.Context) {
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
	router.GET("/json/:name/:who", func(c *gin.Context) {
		name := c.Params.ByName("name")
		switch who := c.Params.ByName("who"); who {
		case "bubble":
			c.String(http.StatusOK, flare.BubbleFlare(name))
		default:
			c.String(http.StatusNotFound, "404 page not found")
		}
	})
}