package main

import (
	"net/http"
	"travek-api/internal/endpoints"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	app := echo.New()
	app.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{http.MethodGet, http.MethodPut, http.MethodPost, http.MethodDelete},
	}))

	endp := endpoints.NewEndpoints()
	app.POST("/country/add", endp.AddCountry)
	app.GET("/country/all", endp.GetAllCountries)
	app.GET("/country", endp.GetOneCountry)
	app.POST("/relation/add", endp.AddRelation)
	app.POST("/road/add", endp.AddRoads)
	app.POST("/tags/add", endp.AddTags)
	app.GET("/tags", endp.GetTags)
	app.POST("/country_tags/add", endp.AddTagsToCountry)
	app.POST("/find", endp.GetSuitable)
	app.Logger.Fatal(app.Start(":8080"))
}
