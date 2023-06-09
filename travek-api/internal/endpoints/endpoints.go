package endpoints

import (
	"fmt"
	"net/http"
	"strconv"
	"travek-api/internal/agragator"
	"travek-api/internal/tags"

	"github.com/labstack/echo/v4"
)

type Endpoints struct{}

func NewEndpoints() Endpoints {
	return Endpoints{}
}

func (*Endpoints) GetAllCountries(c echo.Context) error {
	data, err := agragator.GetAgregator().Cs.GetAllCountries()
	if err != nil {
		return c.JSON(http.StatusOK, BadBaseResponse(err))
	}
	return c.JSON(http.StatusOK, GoodBaseResponse("countries", data))
}

func (*Endpoints) GetOneCountry(c echo.Context) error {
	id, err := strconv.Atoi(c.QueryParam("id"))
	if err != nil {
		return c.JSON(http.StatusOK, BadBaseResponse(err))
	}
	data, err := agragator.GetAgregator().Cs.GetAllCountryDataById(int64(id))
	if err != nil {
		return c.JSON(http.StatusOK, BadBaseResponse(err))
	}
	return c.JSON(http.StatusOK, GoodBaseResponse("country", data))
}
func (*Endpoints) AddCountry(c echo.Context) error {
	request := agragator.BigCountryAddData{}
	err := c.Bind(&request)
	if err != nil {
		return c.JSON(http.StatusOK, BadBaseResponse(err))
	}
	if !request.Validate() {
		return c.JSON(http.StatusOK, BadBaseResponse(NewBaseError("Bad request data")))
	}
	ag := agragator.GetAgregator()
	data, err := ag.AddCountry(&request)
	if err != nil {
		fmt.Println(err.Error())
		return c.JSON(http.StatusOK, BadBaseResponse(err))
	}
	return c.JSON(http.StatusOK, GoodBaseResponse("country", data))
}

func (*Endpoints) AddRelation(c echo.Context) error {
	request := agragator.BigRelationAddData{}
	err := c.Bind(&request)
	if err != nil {
		return c.JSON(http.StatusOK, BadBaseResponse(err))
	}

	ag := agragator.GetAgregator()
	data, err := ag.AddRelations(&request)
	if err != nil {
		fmt.Println(err.Error())
		return c.JSON(http.StatusOK, BadBaseResponse(err))
	}
	return c.JSON(http.StatusOK, GoodBaseResponse("relation", data))
}

func (*Endpoints) AddRoads(c echo.Context) error {
	request := agragator.BigRoadsAddData{}
	err := c.Bind(&request)
	if err != nil {
		return c.JSON(http.StatusOK, BadBaseResponse(err))
	}

	ag := agragator.GetAgregator()
	data, err := ag.AddRoads(&request)
	if err != nil {
		fmt.Println(err.Error())
		return c.JSON(http.StatusOK, BadBaseResponse(err))
	}
	return c.JSON(http.StatusOK, GoodBaseResponse("roads", data))
}

func (*Endpoints) GetTags(c echo.Context) error {
	data, err := agragator.GetAgregator().Tags.GetAllTags()
	if err != nil {
		return c.JSON(http.StatusOK, BadBaseResponse(err))
	}
	return c.JSON(http.StatusOK, GoodBaseResponse("tags", data))
}

func (*Endpoints) AddTags(c echo.Context) error {
	request := []tags.MainTagData{}
	err := c.Bind(&request)
	if err != nil {
		return c.JSON(http.StatusOK, BadBaseResponse(err))
	}

	ag := agragator.GetAgregator()
	data, err := ag.AddTags(&request)
	if err != nil {
		fmt.Println(err.Error())
		return c.JSON(http.StatusOK, BadBaseResponse(err))
	}
	return c.JSON(http.StatusOK, GoodBaseResponse("country_tags", data))
}

func (*Endpoints) AddTagsToCountry(c echo.Context) error {
	request := agragator.TagsToCountryData{}
	err := c.Bind(&request)
	if err != nil {
		return c.JSON(http.StatusOK, BadBaseResponse(err))
	}

	ag := agragator.GetAgregator()
	data, err := ag.AddTagsToCountry(&request)
	if err != nil {
		fmt.Println(err.Error())
		return c.JSON(http.StatusOK, BadBaseResponse(err))
	}
	return c.JSON(http.StatusOK, GoodBaseResponse("country_tags", data))
}

func (*Endpoints) GetSuitable(c echo.Context) error {
	request := agragator.CriteriasData{}
	err := c.Bind(&request)
	if err != nil {
		return c.JSON(http.StatusOK, BadBaseResponse(err))
	}

	ag := agragator.GetAgregator()
	data, err := ag.GetSuitable(&request)
	if err != nil {
		fmt.Println(err.Error())
		return c.JSON(http.StatusOK, BadBaseResponse(err))
	}
	return c.JSON(http.StatusOK, GoodBaseResponse("counties", data))
}
