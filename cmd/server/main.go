package main

import (
	"web/internal/handlers"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.Static("/static", "static")

	e.GET("/", handlers.HomeHandler)
	e.GET("/about", handlers.AboutHandler)
	e.Logger.Fatal(e.Start(":8080"))
}
