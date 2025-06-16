package main

import (
	"github.com/brainboxweb/go-htmx/handlers"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()
	e.Use(middleware.Logger())
	e.GET("", handlers.Home)
	e.GET("/countries", handlers.GetCountries)
	e.GET("/countries/:id", handlers.GetCountry)
	e.GET("/search", handlers.SearchCountries)
	e.Logger.Fatal(e.Start(":42069"))
}
