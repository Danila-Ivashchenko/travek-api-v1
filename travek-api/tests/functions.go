package tests

import "travek-api/internal/country"

func addFakeCountries() []country.PresentedCountryData {
	cs := country.GetService()
	cmdatas := []country.MainCountryData{
		country.MainCountryData{
			Name:      "A",
			Language:  "As",
			Continent: "Eurasia",
		},
		country.MainCountryData{
			Name:      "B",
			Language:  "Bs",
			Continent: "Erope",
		},
		country.MainCountryData{
			Name:      "C",
			Language:  "Cs",
			Continent: "Asia",
		},
		country.MainCountryData{
			Name:      "D",
			Language:  "Ds",
			Continent: "Asia",
		},
	}

	countries := []country.PresentedCountryData{}
	for _, data := range cmdatas {
		result, _ := cs.AddCountry(&data)

		countries = append(countries, *result)
	}
	return countries
}

func deleteFakeCountries(countries []country.PresentedCountryData) error {
	cs := country.GetService()
	for i, data := range countries {
		if i%2 == 0 {
			_, err := cs.DeleteCountryById(data.Id)
			if err != nil {
				return err
			}
		} else {
			_, err := cs.DeleteCountryByName(data.Name)
			if err != nil {
				return err
			}
		}
	}
	return nil
}
