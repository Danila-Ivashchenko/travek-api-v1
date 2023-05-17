package tests

import (
	"testing"
	"travek-api/internal/relation"
)

func TestRelationsIsertDelete(t *testing.T) {
	countries := addFakeCountries()
	defer deleteFakeCountries(countries)

	mdatas := []relation.MainRelationData{
		relation.MainRelationData{
			FirstCountry:  (countries)[0].Id,
			SecondCountry: (countries)[1].Id,
			// FreeEntry:                true,
			PossibilityToStayForever: false,
		},
		relation.MainRelationData{
			FirstCountry:  (countries)[0].Id,
			SecondCountry: (countries)[2].Id,
			// FreeEntry:                true,
			PossibilityToStayForever: false,
		},
		relation.MainRelationData{
			FirstCountry:  (countries)[2].Id,
			SecondCountry: (countries)[3].Id,
			// FreeEntry:                true,
			PossibilityToStayForever: false,
		},
	}
	rs := relation.GetService()

	datas, err := rs.GetAllRelations()
	if err != nil {
		t.Errorf("Error: %s", err.Error())
	}
	relationsLen := len(*datas)

	pdatas := []relation.PresentedRelationData{}
	for _, mdata := range mdatas {
		pdata, err := rs.AddRelation(&mdata)
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
			flag, err = rs.DeleteRelationById(pdata.Id)

		} else {
			flag, err = rs.DeleteRelationByCountriesIds(pdata.FirstCountry, pdata.SecondCountry)
		}
		if !flag || err != nil {
			t.Errorf("Error: %s", err.Error())
		}
	}

	datas, err = rs.GetAllRelations()
	if err != nil {
		t.Errorf("Error: %s", err.Error())
	}

	if len(*datas) != relationsLen {
		t.Errorf("Error: bad get all")
	}

	// for _, pdata := range pdatas {
	// 	countriesData := [2]*country.PresentedCountryData{}
	// 	countriesData[0], err = cs.GetCountyById(pdata.FirstCountry)
	// 	if err != nil {
	// 		t.Errorf("Error: %s", err.Error())
	// 	}
	// 	fmt.Printf("First Country: id = %d, name = %s\n", countriesData[0].Id, countriesData[0].Name)
	// 	countriesData[1], err = cs.GetCountyById(pdata.SecondCountry)
	// 	if err != nil {
	// 		t.Errorf("Error: %s", err.Error())
	// 	}
	// 	fmt.Printf("First Country: id = %d, name = %s\n", countriesData[1].Id, countriesData[1].Name)
	// 	fmt.Printf("Free Entry: %t\n", pdata.FreeEntry)
	// 	fmt.Printf("Possibility ToStay Forever: %t\n", pdata.PossibilityToStayForever)
	// 	fmt.Println()
	// }

}
