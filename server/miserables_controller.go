package server

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"html/template"
)

const (
	TEMPLATE_MISERABLES = TEMPLATES + "miserables.tpl"
)


type MiserablesController struct {}

// This method inits all the routes and deletegates to the associated resources component the gathering of data
func (b *MiserablesController) Run(router *gin.Engine) {

	// miserables diagram page
	router.GET("/miserables", func(c *gin.Context) {
		obj := gin.H{"title": "Miserables Images", "type": "containers"}
		router.SetHTMLTemplate(template.Must(template.ParseFiles(TEMPLATE_MAIN, TEMPLATE_MISERABLES)))
		c.HTML(http.StatusOK, "base", obj)
	})

}