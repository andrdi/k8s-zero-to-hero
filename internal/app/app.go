package app

import (
	"net/http"
	"werf_guide_app/internal/common"
	"werf_guide_app/internal/controllers"

	"github.com/gin-gonic/gin"
)

func Run() {
	route := gin.New()
	route.Use(gin.Recovery())
	route.Use(common.JsonLogger())

	route.Static("/static/stylesheets", "static/stylesheets")
	route.Static("/static/javascripts", "static/javascripts")
	route.Static("/static/images", "static/images")

	route.LoadHTMLGlob("templates/*")

	route.GET("/ping", func(context *gin.Context) {
		context.String(http.StatusOK, "pong\n")
	})

	route.GET("/", controllers.MainPage)
	route.GET("/image", controllers.ImagePage)

	err := route.Run()
	if err != nil {
		return
	}
}
