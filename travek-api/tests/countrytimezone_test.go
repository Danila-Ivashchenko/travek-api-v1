package tests

import (
	"testing"
	"travek-api/internal/country_timezone"
	"travek-api/internal/timezone"
)

func TestCountryTimezone(t *testing.T) {
	countries := addFakeCountries()
	defer deleteFakeCountries(countries)

	ts := timezone.GetService()
	timezones, err := ts.GetAllTimezones()
	if err != nil {
		t.Errorf("Error: %s", err.Error())
	}

	cts := country_timezone.GetService()
	ctdatas, err := cts.GetAllCountryTimezones()
	if err != nil {
		t.Errorf("Error: %s", err.Error())
	}
	ctLen := len(*ctdatas)
	pdatas := []country_timezone.PresentedCountyTimezoneData{}
	for i := range countries {
		ctMainData := country_timezone.MainCountyTimezoneData{
			CountryId: countries[i].Id,
			ZoneId:    (*timezones)[i].Id,
		}
		pdata, err := cts.AddCountryTimezone(&ctMainData)
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
		if i%3 == 0 {
			flag, err = cts.DeleteCountryTimezoneById(pdata.Id)
		} else if i%3 == 1 {
			flag, err = cts.DeleteCountryTimezonesByZoneId(pdata.ZoneId)
		} else {
			flag, err = cts.DeleteCountryTimezonesByCountryId(pdata.CountryId)
		}
		if !flag || err != nil {
			t.Errorf("Error: %s", err.Error())
		}
	}

	ctdatas, err = cts.GetAllCountryTimezones()
	if err != nil {
		t.Errorf("Error: %s", err.Error())
	}
	if ctLen != len(*ctdatas) {
		t.Errorf("Error: %s", "Bad deleting")
	}

}
