package server

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"html/template"
)

const (
	TEMPLATE_BUBBLE = TEMPLATES + "bubble.tpl"
)


type BubbleController struct {}

// This method inits all the routes and deletegates to the associated resources component the gathering of data
func (b *BubbleController) Run(router *gin.Engine) {

	// bubble diagram page
	router.GET("/bubble/:name", func(c *gin.Context) {
		var obj gin.H
		switch name := c.Params.ByName("name"); name {
		case "images":
			obj = gin.H{"title": "Bubble Container", "type": name}
		case "containers":
			obj = gin.H{"title": "Bubble Container", "type": name}
		default:
			c.String(http.StatusNotFound, "404 page not found")
		}
		router.SetHTMLTemplate(template.Must(template.ParseFiles(TEMPLATE_MAIN, TEMPLATE_BUBBLE)))
		c.HTML(http.StatusOK, "base", obj)
	})
}