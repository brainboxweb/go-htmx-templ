package handlers

import (
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
