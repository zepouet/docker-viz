package server

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"html/template"
	"github.com/Treeptik/docker-viz/dockertype"
	"strconv"
)

const (
	TEMPLATE_INDEX = TEMPLATES + "index.tpl"
)


type IndexController struct {}

// This method inits all the routes and deletegates to the associated resources component the gathering of data
func (b *IndexController) Run(router *gin.Engine) {

	// index page
	router.GET("/", func(c *gin.Context) {
		countImg := dockertype.CountDockerImages()
		countCont := dockertype.CountDockerContainer()
		obj := gin.H{"title": "Dashboard", "countImages": strconv.Itoa(countImg), "countContainers": strconv.Itoa(countCont)}
		router.SetHTMLTemplate(template.Must(template.ParseFiles(TEMPLATE_MAIN, TEMPLATE_INDEX)))
		c.HTML(http.StatusOK, "base", obj)
	})
}