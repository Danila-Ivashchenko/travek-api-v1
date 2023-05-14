package main

import (
	"travek-api/internal/endpoints"

	"github.com/labstack/echo/v4"
)

func main() {
	app := echo.New()
	endp := endpoints.NewEndpoints()
	app.POST("/country/add", endp.AddCountry)
	app.POST("/relation/add", endp.AddRelation)
	app.POST("/road/add", endp.AddRoads)
	app.POST("/tags/add", endp.AddTags)
	app.POST("/country_tags/add", endp.AddTagsToCountry)
	app.POST("/find", endp.GetSuitable)
	app.Logger.Fatal(app.Start(":8080"))
}
