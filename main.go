package main

import (
	"ginWeb/routes"
	"github.com/gin-gonic/gin"
)

func main() {

	var app = gin.Default()

	routes.InitRoutes(app, routes.Routes)

	app.LoadHTMLGlob("./tem/*")
	
	startServer(app, ":9091")
}

func startServer(app *gin.Engine, port string) {
	app.Run(port) // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
