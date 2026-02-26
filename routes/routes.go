package routes

import (
	_ "apis/docs"
	"github.com/labstack/echo/v5"
	echoSwagger "github.com/swaggo/echo-swagger"
)

// @title APIs
// @version 0.1
// @description Documentation for all the apis
// @host localhost:3000
// @BasePath /api
func Routes(e *echo.Echo) {
	e.GET("/swagger/*", echoSwagger.WrapHandler)

	base := e.Group("/api/v0.1")
	registerCompetitionRoutes(base)
}
