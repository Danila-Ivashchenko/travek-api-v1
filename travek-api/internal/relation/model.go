package relation

import "fmt"

type MainRelationData struct {
	FirstCountry             int64 `json:"first_country"`
	SecondCountry            int64 `json:"second_country"`
	FreeEntry                bool  `json:"free_entry"`
	PossibilityToStayForever bool  `json:"possibility_to_stay_forever"`
}

func (m *MainRelationData) sqlInsertString() string {
	return fmt.Sprintf(`INSERT INTO relation (first_country, second_country, free_entry, possibility_to_stay_forever) VALUES %s`, m.sqlValuesString())
}

func (m *MainRelationData) sqlValuesString() string {
	return fmt.Sprintf(`(%d, %d, %t, %t)`, m.FirstCountry, m.SecondCountry, m.FreeEntry, m.PossibilityToStayForever)
}

type PresentedRelationData struct {
	Id int64 `json:"id"`
	MainRelationData
}

func PresentedRelationDataFromMain(data *MainRelationData, id int64) PresentedRelationData {
	return PresentedRelationData{Id: id, MainRelationData: *data}
}

// errors

type relationExistError struct{}

func (*relationExistError) Error() string {
	return "This relation is already exist"
}

type relationNotExistError struct{}

func (*relationNotExistError) Error() string {
	return "This relation is not exist"
}
