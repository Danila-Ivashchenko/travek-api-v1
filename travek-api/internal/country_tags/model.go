package country_tags

import "fmt"

type MainCountyTagsData struct {
	CountryId int64 `json:"country_id"`
	TagId int64 `json:"tag_id"`
}

func (m *MainCountyTagsData) sqlInsertString() string {
	return fmt.Sprintf(`INSERT INTO country_tags (country_id, tag_id) VALUES %s`, m.sqlValuesString())
}

func (m *MainCountyTagsData) sqlValuesString() string {
	return fmt.Sprintf(`(%d, %d)`,m.CountryId, m.TagId)
}

type PresentedCountyTagsData struct {
	Id int64 `json:"id"`
	MainCountyTagsData
}

func PresentedCountyTagsDataFromMain(data *MainCountyTagsData, id int64) PresentedCountyTagsData {
	return PresentedCountyTagsData{Id: id, MainCountyTagsData: *data}
}

// errors

type country_tagExistError struct{}

func (*country_tagExistError) Error() string {
	return "This country_tag is already exist"
}
