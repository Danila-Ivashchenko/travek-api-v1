package timezone

import "fmt"

type MainTimezoneData struct {
	Zone int `json:"zone"`
}

func (m *MainTimezoneData) sqlInsertString() string {
	return fmt.Sprintf(`INSERT INTO timezone (zone) VALUES %s`, m.sqlValuesString())
}

func (m *MainTimezoneData) sqlValuesString() string {
	return fmt.Sprintf(`(%d)`, m.Zone)
}

type PresentedTimezoneData struct {
	Id int64 `json:"id"`
	MainTimezoneData
}

func PresentedTimezoneDataFromMain(data *MainTimezoneData, id int64) PresentedTimezoneData {
	return PresentedTimezoneData{Id: id, MainTimezoneData: *data}
}

// errors

type timezoneExistError struct{}

func (*timezoneExistError) Error() string {
	return "This timezone is already exist"
}
