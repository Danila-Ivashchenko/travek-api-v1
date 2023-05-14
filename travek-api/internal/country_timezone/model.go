package country_timezone

import "fmt"

type MainCountyTimezoneData struct {
	CountryId int64 `json:"country_id"`
	ZoneId int64 `json:"zone_id"`
}

func (m *MainCountyTimezoneData) sqlInsertString() string {
	return fmt.Sprintf(`INSERT INTO country_timezones (country_id, timezone_id) VALUES %s`, m.sqlValuesString())
}

func (m *MainCountyTimezoneData) sqlValuesString() string {
	return fmt.Sprintf(`(%d, %d)`,m.CountryId, m.ZoneId)
}

type PresentedCountyTimezoneData struct {
	Id int64 `json:"id"`
	MainCountyTimezoneData
}

func PresentedCountyTimezoneDataFromMain(data *MainCountyTimezoneData, id int64) PresentedCountyTimezoneData {
	return PresentedCountyTimezoneData{Id: id, MainCountyTimezoneData: *data}
}

// errors

type country_timezoneExistError struct{}

func (*country_timezoneExistError) Error() string {
	return "This country_timezone is already exist"
}
