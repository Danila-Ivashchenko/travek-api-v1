package country_timezone

import (
	"fmt"
	"travek-api/internal/timezone"
	"travek-api/pkg/database"
)

type rowScaner interface {
	Scan(dest ...any) error
}

func presentedDataFromRowScaner(row rowScaner) (*PresentedCountyTimezoneData, error) {
	data := &PresentedCountyTimezoneData{}
	err := row.Scan(&data.Id, &data.CountryId, &data.ZoneId)
	if err != nil {
		return nil, err
	}
	return data, nil
}

type Servise interface {
	AddCountryTimezone(data *MainCountyTimezoneData) (*PresentedCountyTimezoneData, error)
	AddCountryTimezones(country_id int64, timezones *[]timezone.PresentedTimezoneData) (int, error)

	GetCountryTimezoneById(id int64) (*PresentedCountyTimezoneData, error)
	GetCountryTimezoneByCountryAndZoneIds(country_id, timezone_id int64) (*PresentedCountyTimezoneData, error)

	GetAllCountryTimezones(limit_ofset ...int) (*[]PresentedCountyTimezoneData, error)
	GetCountryTimezonesByCountryId(id int64, limit_ofset ...int) (*[]PresentedCountyTimezoneData, error)
	GetCountryTimezonesByZoneId(id int64, limit_ofset ...int) (*[]PresentedCountyTimezoneData, error)

	DeleteCountryTimezoneById(id int64) (bool, error)
	DeleteCountryTimezonesByZoneId(id int64) (bool, error)
	DeleteCountryTimezonesByCountryId(id int64) (bool, error)
}

type servise struct{}

// add

func (s *servise) AddCountryTimezone(data *MainCountyTimezoneData) (*PresentedCountyTimezoneData, error) {
	if s.checkCountryTimezoneExist(data.CountryId, data.ZoneId) {
		return nil, &country_timezoneExistError{}
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
	pdata := PresentedCountyTimezoneDataFromMain(data, id)
	return &pdata, nil
}

func (s *servise) AddCountryTimezones(country_id int64, timezones *[]timezone.PresentedTimezoneData) (int, error) {
	stmt := "INSERT INTO country_timezones (country_id, timezone_id) VALUES "
	for i, timezoneData := range *timezones {
		comma := ""
		if i > 0 {
			comma = ", "
		}
		stmt += fmt.Sprintf("%s(%d, %d) ", comma, country_id, timezoneData.Id)
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

func (s *servise) checkCountryTimezoneExist(country_id, timezone_id int64) bool {
	data, _ := s.GetCountryTimezoneByCountryAndZoneIds(country_id, timezone_id)
	return data != nil
}

// primar selects

func (*servise) getRowWithWhereCase(whereCase string) (*PresentedCountyTimezoneData, error) {
	stmt := "SELECT * FROM country_timezones"
	if whereCase != "" {
		stmt += fmt.Sprintf("WHERE %s", whereCase)
	}
	db := database.Get_db()
	defer db.Close()

	row := db.QueryRow(stmt)
	data, err := presentedDataFromRowScaner(row)
	return data, err
}

func (*servise) getRowsWithWhereCase(whereCase string, limit_ofset ...int) (*[]PresentedCountyTimezoneData, error) {
	stmt := "SELECT * FROM country_timezones"
	if whereCase != "" {
		stmt += fmt.Sprintf("WHERE %s", whereCase)
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
	pdatas := []PresentedCountyTimezoneData{}
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

func (s *servise) GetCountryTimezoneById(id int64) (*PresentedCountyTimezoneData, error) {
	whereCase := fmt.Sprintf(`country_timezones WHERE id = %d`, id)
	return s.getRowWithWhereCase(whereCase)
}

func (s *servise) GetCountryTimezoneByCountryAndZoneIds(country_id, timezone_id int64) (*PresentedCountyTimezoneData, error) {
	whereCase := fmt.Sprintf(`country_id = %d AND timezone_id = %d`, country_id, timezone_id)
	return s.getRowWithWhereCase(whereCase)
}

// select a lot

func (s *servise) GetAllCountryTimezones(limit_ofset ...int) (*[]PresentedCountyTimezoneData, error) {
	return s.getRowsWithWhereCase("", limit_ofset...)
}

func (s *servise) GetCountryTimezonesByZoneId(timezone_id int64, limit_ofset ...int) (*[]PresentedCountyTimezoneData, error) {
	whereCase := fmt.Sprintf("timezone_id = %d", timezone_id)
	return s.getRowsWithWhereCase(whereCase, limit_ofset...)
}

func (s *servise) GetCountryTimezonesByCountryId(country_id int64, limit_ofset ...int) (*[]PresentedCountyTimezoneData, error) {
	whereCase := fmt.Sprintf("country_id = %d", country_id)
	return s.getRowsWithWhereCase(whereCase, limit_ofset...)
}

//

func (*servise) DeleteCountryTimezoneById(id int64) (bool, error) {
	stmt := fmt.Sprintf(`DELETE FROM country_timezones WHERE id = %d LIMIT 1`, id)
	db := database.Get_db()
	defer db.Close()

	_, err := db.Exec(stmt)
	if err != nil {
		return false, err
	}
	return true, err
}

func (*servise) DeleteCountryTimezonesByZoneId(timezone_id int64) (bool, error) {
	stmt := fmt.Sprintf(`DELETE FROM country_timezones WHERE timezone_id = %d`, timezone_id)
	db := database.Get_db()
	defer db.Close()

	_, err := db.Exec(stmt)
	if err != nil {
		return false, err
	}
	return true, err
}

func (*servise) DeleteCountryTimezonesByCountryId(country_id int64) (bool, error) {
	stmt := fmt.Sprintf(`DELETE FROM country_timezones WHERE country_id = %d`, country_id)
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
