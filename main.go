package main

import (
	"apis/routes"

	"log"

	"github.com/labstack/echo/v5"
)

func main() {
	e := echo.New()

	// Add Routes
	routes.Routes(e)

	// Start server
	if err := e.Start(":3000"); err != nil {
		log.Fatal(err)
	}
}
