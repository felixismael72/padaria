package router

import "github.com/labstack/echo/v4"

func Start() *echo.Echo {
	server := echo.New()

	api := server.Group("/api")

	loadProductRoutes(api)

	return server
}
