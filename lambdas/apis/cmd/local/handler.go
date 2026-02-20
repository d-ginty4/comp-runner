package local

import (
	"apis/routes"
	"github.com/labstack/echo/v5"
)

func Handler() {
	e := echo.New()
	routes.Routes(e)

	if err := e.Start(":3000"); err != nil {
		panic(err)
	}
}
