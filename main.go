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

	// Countries
	e.GET("/countries", handlers.GetCountries)
	e.GET("/countries/:id", handlers.GetCountry)
	e.GET("/search", handlers.SearchCountries)

	// Contacts
	e.GET("/contacts", handlers.GetContacts)
	e.POST("/contacts", handlers.PostContacts)  
		
	e.Logger.Fatal(e.Start(":42069"))
}
