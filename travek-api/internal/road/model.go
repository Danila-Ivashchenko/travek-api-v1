package road

import "fmt"

type MainRoadData struct {
	FirstCountry  int64  `json:"first_country"`
	SecondCountry int64  `json:"second_country"`
	Transport     string `json:"transport"`
	TimeHours     int    `json:"time_hours"`
}

func (m *MainRoadData) sqlInsertString() string {
	return fmt.Sprintf(`INSERT INTO road (first_country, second_country, transport, time_hours) VALUES %s`, m.sqlValuesString())
}

func (m *MainRoadData) sqlValuesString() string {
	return fmt.Sprintf(`(%d, %d, "%s", %d)`, m.FirstCountry, m.SecondCountry, m.Transport, m.TimeHours)
}

type PresentedRoadData struct {
	Id int64 `json:"id"`
	MainRoadData
}

func PresentedRoadDataFromMain(data *MainRoadData, id int64) PresentedRoadData {
	return PresentedRoadData{Id: id, MainRoadData: *data}
}

// errors

type roadExistError struct{}

func (*roadExistError) Error() string {
	return "This road is already exist"
}

// type roadNotExistError struct{}

// func (*roadNotExistError) Error() string {
// 	return "This road is not exist"
// }
