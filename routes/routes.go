package routes

import (
	"ginWeb/services"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

type Route struct {
	Method     string
	Url        string
	Controller func(ctx *gin.Context)
	Children   []Route
}

type AppOrSubApp interface {
	GET(relativePath string, handlers ...gin.HandlerFunc) gin.IRoutes
	POST(relativePath string, handlers ...gin.HandlerFunc) gin.IRoutes
	PUT(relativePath string, handlers ...gin.HandlerFunc) gin.IRoutes
	PATCH(relativePath string, handlers ...gin.HandlerFunc) gin.IRoutes
	DELETE(relativePath string, handlers ...gin.HandlerFunc) gin.IRoutes
	Group(relativePath string, handlers ...gin.HandlerFunc) *gin.RouterGroup
}

var Routes []Route = []Route{
	{
		Method:     http.MethodGet,
		Url:        "/product",
		Controller: services.PingController,
	},
	{
		Method:     http.MethodGet,
		Url:        "/time",
		Controller: services.TimerController,
	},
	{
		Method:     http.MethodGet,
		Url:        "/login",
		Controller: services.LoginController,
	},
	{
		Method:     http.MethodPost,
		Url:        "/auth",
		Controller: services.AuthController,
	},
	{
		Method:     http.MethodGet,
		Url:        "/hello",
		Controller: services.HelloController,
	},
	OrderRoutes,
}

func InitRoutes(app AppOrSubApp, routes []Route) {

	for _, route := range routes {

		if len(route.Children) == 0 {

			var handler func(relativePath string, handlers ...gin.HandlerFunc) gin.IRoutes

			switch strings.ToUpper(route.Method) {

			case http.MethodPut:
				handler = app.PUT

			case http.MethodPost:
				handler = app.POST

			case http.MethodDelete:
				handler = app.DELETE

			case http.MethodPatch:
				handler = app.PATCH

			default:
				handler = app.GET

			}

			handler(route.Url, route.Controller)

		} else {

			var group = app.Group(route.Url)

			InitRoutes(group, route.Children)

		}

	}

}
