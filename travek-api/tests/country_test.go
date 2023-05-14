package tests

import (
	"testing"
	"travek-api/internal/country"
)

func TestCountriesIsertDelete(t *testing.T) {
	// mdatas := []country.MainCountryData{
	// 	country.MainCountryData{
	// 		Name:      "Russia",
	// 		Language:  "Russian",
	// 		Continent: "Eurasia",
	// 	},
	// 	country.MainCountryData{
	// 		Name:      "USA",
	// 		Language:  "English",
	// 		Continent: "North America",
	// 	},
	// 	country.MainCountryData{
	// 		Name:      "Japan",
	// 		Language:  "Japanese",
	// 		Continent: "Asia",
	// 	},
	// 	country.MainCountryData{
	// 		Name:      "Canada",
	// 		Language:  "English",
	// 		Continent: "North America",
	// 	},
	// 	country.MainCountryData{
	// 		Name:      "Germany",
	// 		Language:  "German",
	// 		Continent: "Europe",
	// 	},
	// 	country.MainCountryData{
	// 		Name:      "Australia",
	// 		Language:  "English",
	// 		Continent: "Australia",
	// 	},
	// 	country.MainCountryData{
	// 		Name:      "Chine",
	// 		Language:  "Chinese",
	// 		Continent: "Asia",
	// 	},
	// 	country.MainCountryData{
	// 		Name:      "Greate Britain",
	// 		Language:  "English",
	// 		Continent: "Europe",
	// 	},
	// 	country.MainCountryData{
	// 		Name:      "Singapore",
	// 		Language:  "English",
	// 		Continent: "Asia",
	// 	},
	// 	country.MainCountryData{
	// 		Name:      "France",
	// 		Language:  "Franch",
	// 		Continent: "Europe",
	// 	},
	// }

	mdatas := []country.MainCountryData{
		country.MainCountryData{
			Name:      "A",
			Language:  "As",
			Continent: "Eurasia",
		},
		country.MainCountryData{
			Name:      "B",
			Language:  "Bs",
			Continent: "Eurasia",
		},
	}

	s := country.GetService()
	pdatas := []country.PresentedCountryData{}
	for _, data := range mdatas {
		result, err := s.AddCountry(&data)
		if err != nil {
			t.Errorf("Error: %s", err.Error())
		}
		pdatas = append(pdatas, *result)
	}

	for i, data := range pdatas {
		if i%2 == 0 {
			_, err := s.DeleteCountryById(data.Id)
			if err != nil {
				t.Errorf("Error: %s", err.Error())
			}
		} else {
			_, err := s.DeleteCountryByName(data.Name)
			if err != nil {
				t.Errorf("Error: %s", err.Error())
			}
		}
	}
}
