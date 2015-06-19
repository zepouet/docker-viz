package server

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"github.com/Treeptik/docker-viz/dockertype"
)

type DockerController struct {}

/* This method inits all the routes and deletegates to the associated resources component the gathering of data */
func (b *DockerController) Run(router *gin.Engine) {

	// docker success connexion
	router.GET("/docker", func(c *gin.Context) {
		if dockertype.DockerStatut() {
			c.String(http.StatusOK, "true")
		} else {
			c.String(http.StatusOK, "false")
		}
	})
}