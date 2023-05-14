package country_tags

import (
	"fmt"
	"travek-api/internal/tags"
	"travek-api/pkg/database"
)

type rowScaner interface {
	Scan(dest ...any) error
}

func presentedDataFromRowScaner(row rowScaner) (*PresentedCountyTagsData, error) {
	data := &PresentedCountyTagsData{}
	err := row.Scan(&data.Id, &data.CountryId, &data.TagId)
	if err != nil {
		return nil, err
	}
	return data, nil
}

type Servise interface {
	AddCountyTag(data *MainCountyTagsData) (*PresentedCountyTagsData, error)
	AddCountyTags(country_id int64, tags *[]tags.PresentedTagData) (int, error)

	GetCountyTagById(id int64) (*PresentedCountyTagsData, error)
	//GetCountyTagByCountryAndTagIds(country_id, tag_id int64) (*PresentedCountyTagsData, error)

	GetAllCountyTags(limit_ofset ...int) (*[]PresentedCountyTagsData, error)
	GetCountyTagsByCountryId(id int64, limit_ofset ...int) (*[]PresentedCountyTagsData, error)
	GetCountyTagsByTagId(id int64, limit_ofset ...int) (*[]PresentedCountyTagsData, error)

	DeleteCountyTagsById(id int64) (bool, error)
	DeleteCountyTagsByTagId(id int64) (bool, error)
	DeleteCountyTagsByCountryId(id int64) (bool, error)
}

type servise struct{}

// add

func (s *servise) AddCountyTag(data *MainCountyTagsData) (*PresentedCountyTagsData, error) {
	if s.checkCountyTagExist(data.CountryId, data.TagId) {
		return nil, &country_tagExistError{}
	}
	db := database.Get_db()
	defer db.Close()

	result, err := db.Exec(data.sqlInsertString())
	if err != nil {
		return nil, err
	}
	id, err := result.LastInsertId()
	if err != nil {
		return nil, err
	}
	pdata := PresentedCountyTagsDataFromMain(data, id)
	return &pdata, nil
}

func (s *servise) AddCountyTags(country_id int64, tags *[]tags.PresentedTagData) (int, error) {
	stmt := "INSERT INTO country_tags (country_id, tag_id) VALUES "
	for _, tag := range *tags {
		if s.checkCountyTagExist(country_id, tag.Id) {
			return -1, &country_tagExistError{}
		}
	}
	for i, tagData := range *tags {
		comma := ""
		if i > 0 {
			comma = ", "
		}
		stmt += fmt.Sprintf("%s(%d, %d) ", comma, country_id, tagData.Id)
	}

	db := database.Get_db()
	defer db.Close()

	result, err := db.Exec(stmt)
	if err != nil {
		return -1, err
	}
	count, err := result.RowsAffected()
	if err != nil {
		return -1, err
	}

	return int(count), nil
}

// check

func (s *servise) checkCountyTagExist(country_id, tag_id int64) bool {
	data, _ := s.GetCountyTagByCountryAndZoneIds(country_id, tag_id)
	return data != nil
}

// primar selects

func (*servise) getRowWithWhereCase(whereCase string) (*PresentedCountyTagsData, error) {
	stmt := "SELECT * FROM country_tags"
	if whereCase != "" {
		stmt += fmt.Sprintf("WHERE %s", whereCase)
	}
	db := database.Get_db()
	defer db.Close()

	row := db.QueryRow(stmt)
	data, err := presentedDataFromRowScaner(row)
	return data, err
}

func (*servise) getRowsWithWhereCase(whereCase string, limit_ofset ...int) (*[]PresentedCountyTagsData, error) {
	stmt := "SELECT * FROM country_tags"
	if whereCase != "" {
		stmt += fmt.Sprintf(" WHERE %s", whereCase)
	}
	if len(limit_ofset) >= 1 {
		stmt += fmt.Sprintf(` LIMIT %d`, limit_ofset[0])
	}
	if len(limit_ofset) >= 2 {
		stmt += fmt.Sprintf(`, %d`, limit_ofset[1])
	}

	db := database.Get_db()
	defer db.Close()

	rows, err := db.Query(stmt)
	if err != nil {
		return nil, err
	}
	pdatas := []PresentedCountyTagsData{}
	for rows.Next() {
		pdata, err := presentedDataFromRowScaner(rows)
		if err != nil {
			break
		}
		pdatas = append(pdatas, *pdata)
	}
	return &pdatas, nil
}

// select one

func (s *servise) GetCountyTagById(id int64) (*PresentedCountyTagsData, error) {
	whereCase := fmt.Sprintf(`country_tags WHERE id = %d`, id)
	return s.getRowWithWhereCase(whereCase)
}

func (s *servise) GetCountyTagByCountryAndZoneIds(country_id, tag_id int64) (*PresentedCountyTagsData, error) {
	whereCase := fmt.Sprintf(`country_id = %d AND tag_id = %d`, country_id, tag_id)
	return s.getRowWithWhereCase(whereCase)
}

// select a lot

func (s *servise) GetAllCountyTags(limit_ofset ...int) (*[]PresentedCountyTagsData, error) {
	return s.getRowsWithWhereCase("", limit_ofset...)
}

func (s *servise) GetCountyTagsByTagId(tag_id int64, limit_ofset ...int) (*[]PresentedCountyTagsData, error) {
	whereCase := fmt.Sprintf("tag_id = %d", tag_id)
	return s.getRowsWithWhereCase(whereCase, limit_ofset...)
}

func (s *servise) GetCountyTagsByCountryId(country_id int64, limit_ofset ...int) (*[]PresentedCountyTagsData, error) {
	whereCase := fmt.Sprintf("country_id = %d", country_id)
	return s.getRowsWithWhereCase(whereCase, limit_ofset...)
}

//

func (*servise) DeleteCountyTagsById(id int64) (bool, error) {
	stmt := fmt.Sprintf(`DELETE FROM country_tags WHERE id = %d LIMIT 1`, id)
	db := database.Get_db()
	defer db.Close()

	_, err := db.Exec(stmt)
	if err != nil {
		return false, err
	}
	return true, err
}

func (*servise) DeleteCountyTagsByTagId(tag_id int64) (bool, error) {
	stmt := fmt.Sprintf(`DELETE FROM country_tags WHERE tag_id = %d`, tag_id)
	db := database.Get_db()
	defer db.Close()

	_, err := db.Exec(stmt)
	if err != nil {
		return false, err
	}
	return true, err
}

func (*servise) DeleteCountyTagsByCountryId(country_id int64) (bool, error) {
	stmt := fmt.Sprintf(`DELETE FROM country_tags WHERE country_id = %d`, country_id)
	db := database.Get_db()
	defer db.Close()

	_, err := db.Exec(stmt)
	if err != nil {
		return false, err
	}
	return true, err
}

func GetService() Servise {
	return &servise{}
}
