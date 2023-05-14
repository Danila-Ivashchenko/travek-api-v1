package tests

import (
	"testing"
	"travek-api/internal/road"
)


func TestRoadIsertDelete(t *testing.T) {
	countries := addFakeCountries()
	defer deleteFakeCountries(countries)

	mrdatas := []road.MainRoadData{
		road.MainRoadData{
			FirstCountry:  countries[0].Id,
			SecondCountry: countries[1].Id,
			Transport:     "train",
			TimeHours:     12,
		},
		road.MainRoadData{
			FirstCountry:  countries[0].Id,
			SecondCountry: countries[2].Id,
			Transport:     "plain",
			TimeHours:     6,
		},
		road.MainRoadData{
			FirstCountry:  countries[0].Id,
			SecondCountry: countries[3].Id,
			Transport:     "train",
			TimeHours:     2,
		},
		road.MainRoadData{
			FirstCountry:  countries[1].Id,
			SecondCountry: countries[3].Id,
			Transport:     "plain",
			TimeHours:     1,
		},
	}
	rs := road.GetService()
	roads, err := rs.GetAllRoads()
	if err != nil {
		t.Errorf("Error: %s", err.Error())
	}
	roadsLen := len(*roads)

	pdatas := []road.PresentedRoadData{}
	for _, data := range mrdatas {
		pdata, err := rs.AddRoad(&data)
		if err != nil {
			t.Errorf("Error: %s", err.Error())
		}
		pdatas = append(pdatas, *pdata)
	}

	for i, pdata := range pdatas {
		var (
			flag bool
			err  error
		)
		if i%2 == 0 {
			flag, err = rs.DeleteRoadById(pdata.Id)

		} else {
			flag, err = rs.DeleteRoadByCountriesIds(pdata.FirstCountry, pdata.SecondCountry)
		}
		if !flag || err != nil {
			t.Errorf("Error: %s", err.Error())
		}
	}

	newRoads, err := rs.GetAllRoads()
	if err != nil {
		t.Errorf("Error: %s", err.Error())
	}
	if len(*newRoads) != roadsLen {
		t.Errorf("Error: %s", "Bad get all")
	}

}
