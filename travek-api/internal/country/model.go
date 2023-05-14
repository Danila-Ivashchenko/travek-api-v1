package country

import "fmt"

type MainCountryData struct {
	Name      string `json:"name"`
	Language  string `json:"language"`
	Continent string `json:"continent"`
}

func (m *MainCountryData) sqlInsertString() string {
	return fmt.Sprintf(`INSERT INTO country (name, language, continent) VALUES %s`, m.sqlValuesString())
}

func (m *MainCountryData) sqlValuesString() string {
	return fmt.Sprintf(`("%s", "%s", "%s")`, m.Name, m.Language, m.Continent)
}

type PresentedCountryData struct {
	Id int64 `json:"id"`
	MainCountryData
}

func PresentedCountryDataFromMain(data *MainCountryData, id int64) PresentedCountryData {
	return PresentedCountryData{Id: id, MainCountryData: *data}
}

// errors

type countryExistError struct{}

func (*countryExistError) Error() string {
	return "This country is already exist"
}

type countryNotExistError struct{}

func (*countryNotExistError) Error() string {
	return "This country is not exist"
}