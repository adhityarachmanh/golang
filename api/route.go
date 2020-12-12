package api

import (
	"github.com/gin-gonic/gin"
)

func (app *AppSchema) initializeRoutes() {
	// merge semua route
	app.Router = gin.Default()
	app.Router.Use(gin.Logger())
	app.Router.LoadHTMLGlob("templates/*.html")
	app.Router.Static("/static", "static")
	app.Router.GET("/", app.index)
	app.routeRegister("POST", "visitor/information", app.visitorIP, false)
}
