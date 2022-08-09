package route

import (
	"github.com/agusbasari29/simple-xjx-tasking.git/helper"
	"github.com/gin-gonic/gin"
)

func DefineApiRoute(e *gin.Engine) {
	handlers := []helper.Handler{
		TaskRoutes{},
	}
	var routes []helper.Route
	for _, handler := range handlers {
		routes = append(routes, handler.Route()...)
	}
	api := e.Group("/api/v1")
	for _, route := range routes {
		switch route.Method {
		case "GET":
			{
				api.GET(route.Path, route.Handler...)
			}
		case "POST":
			{
				api.POST(route.Path, route.Handler...)
			}
		case "PUT":
			{
				api.PUT(route.Path, route.Handler...)
			}
		case "PATCH":
			{
				api.PATCH(route.Path, route.Handler...)
			}
		case "DELETE":
			{
				api.DELETE(route.Path, route.Handler...)
			}
		}
	}
}
