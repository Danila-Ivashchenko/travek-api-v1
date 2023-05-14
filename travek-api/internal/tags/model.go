package tags

import "fmt"

type MainTagData struct {
	Name string `json:"name"`
}

func (m *MainTagData) sqlInsertString() string {
	return fmt.Sprintf(`INSERT INTO tags (name) VALUES %s`, m.sqlValuesString())
}

func (m *MainTagData) sqlValuesString() string {
	return fmt.Sprintf(`("%s")`, m.Name)
}

type PresentedTagData struct {
	Id int64 `json:"id"`
	MainTagData
}

func PresentedTagDataFromMain(data *MainTagData, id int64) PresentedTagData {
	return PresentedTagData{Id: id, MainTagData: *data}
}

// errors

type tagExistError struct{}

func (*tagExistError) Error() string {
	return "This tag is already exist"
}
