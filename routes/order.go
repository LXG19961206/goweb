package routes

import (
	"ginWeb/services"
	"net/http"
)

var OrderRoutes = Route{
	Url: "/order",
	Children: []Route{
		{
			Method:     http.MethodPost,
			Url:        "/list",
			Controller: services.PingController,
		},
		{
			Method:     http.MethodGet,
			Url:        "/detail",
			Controller: services.HelloController,
		},
	},
}
