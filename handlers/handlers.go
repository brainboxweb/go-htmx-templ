package handlers

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/brainboxweb/go-htmx/components"
	"github.com/brainboxweb/go-htmx/models"
	"github.com/labstack/echo/v4"
)

func Home(ctx echo.Context) error {
	component := components.Home()
	return component.Render(ctx.Request().Context(), ctx.Response().Writer)
}

func GetCountries(ctx echo.Context) error {
	component := components.Countries(models.Countries)
	return component.Render(ctx.Request().Context(), ctx.Response().Writer)
}

func GetCountry(ctx echo.Context) error {
	countryName := ctx.Param("id")
	var required models.Country
	for _, country := range models.Countries {
		if country.Name == countryName {
			required = country
			component := components.Country(required)
			return component.Render(ctx.Request().Context(), ctx.Response().Writer)
		}
	}
	return echo.NewHTTPError(http.StatusNotFound, "Country not found")
}

func SearchCountries(ctx echo.Context) error {
	query := strings.ToLower(ctx.QueryParam("search"))
	var results []models.Country
	for _, country := range models.Countries {
		if strings.Contains(strings.ToLower(country.Name), query) {
			results = append(results, country)
		}
	}
	component := components.CountryList(results)
	return component.Render(ctx.Request().Context(), ctx.Response().Writer)
}

// Contacts

var data = models.NewData() // In-memory data store

func GetContacts(ctx echo.Context) error {
	pg := models.Page{Data: data, Form: models.FormData{}}
	component := components.GetContacts(pg)
	return component.Render(ctx.Request().Context(), ctx.Response().Writer)
}

func PostContacts(ctx echo.Context) error {

	name := ctx.FormValue("name")
	email := ctx.FormValue("email")



	// check for dupes
	if data.HasEmail(email) {
		formData := models.NewFormData()
		formData.Values["name"] = name
		formData.Values["email"] = email

		fmt.Println("It s a DUP!!!")
		formData.Errors["email"] = "Dupe!"

		// 	return c.Render(422, "form", formData)
		writer := ctx.Response().Writer
		writer.WriteHeader(http.StatusUnprocessableEntity) // 422

		pg := models.Page{Data: data, Form: formData}
		component := components.JustContacts(pg)
		return component.Render(ctx.Request().Context(), writer)
	}

	// Validate
	if !(name == "" && email == "") {
		data.Contacts = append(data.Contacts, models.NewContact(name, email))
	}
	
	pg := models.Page{Data: data, Form: models.FormData{}}
	component := components.JustContacts(pg)
	return component.Render(ctx.Request().Context(), ctx.Response().Writer)
}
