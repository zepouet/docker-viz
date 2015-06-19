package server

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"html/template"
)

const (
	TEMPLATE_DENDROGAM = TEMPLATES + "dendrogam.tpl"
)


type DendrogramController struct {}

// This method inits all the routes and deletegates to the associated resources component the gathering of data
func (b *DendrogramController) Run(router *gin.Engine) {

	// dendrogam diagram page
	router.GET("/dendrogam", func(c *gin.Context) {
		obj := gin.H{"title": "Dendrogam Images", "type": "images"}
		router.SetHTMLTemplate(template.Must(template.ParseFiles(TEMPLATE_MAIN, TEMPLATE_DENDROGAM)))
		c.HTML(http.StatusOK, "base", obj)
	})
}