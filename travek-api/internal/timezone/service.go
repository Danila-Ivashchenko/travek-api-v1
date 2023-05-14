package timezone

import (
	"fmt"
	"strconv"
	"strings"
	"travek-api/pkg/database"
)

type rowScaner interface {
	Scan(dest ...any) error
}

func presentedDataFromRowScaner(row rowScaner) (*PresentedTimezoneData, error) {
	data := &PresentedTimezoneData{}
	err := row.Scan(&data.Id, &data.Zone)
	if err != nil {
		return nil, err
	}
	return data, nil
}

type Servise interface {
	AddTimezone(data *MainTimezoneData) (*PresentedTimezoneData, error)

	GetSeveralTimezones(zones []int, limit_ofset ...int) (*[]PresentedTimezoneData, error)
	GetAllTimezones(limit_ofset ...int) (*[]PresentedTimezoneData, error)

	GetTimezoneById(id int64) (*PresentedTimezoneData, error)
	GetTimezoneByZone(id int) (*PresentedTimezoneData, error)

	DeleteTimezoneById(id int64) (bool, error)
	DeleteTimezoneByZone(zone int) (bool, error)
}

type servise struct{}

// primal select

func (*servise) getRowWithWhereCase(whereCase string) (*PresentedTimezoneData, error) {
	stmt := "SELECT * FROM timezone"
	if whereCase != "" {
		stmt += fmt.Sprintf(" WHERE %s", whereCase)
	}
	db := database.Get_db()
	defer db.Close()

	row := db.QueryRow(stmt)
	data, err := presentedDataFromRowScaner(row)
	return data, err
}

func (*servise) getRowsWithWhereCase(whereCase string, limit_ofset ...int) (*[]PresentedTimezoneData, error) {
	stmt := "SELECT * FROM timezone"
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
	pdatas := []PresentedTimezoneData{}
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

func (s *servise) GetTimezoneById(id int64) (*PresentedTimezoneData, error) {
	whereCase := fmt.Sprintf("id = %d", id)
	return s.getRowWithWhereCase(whereCase)
}

func (s *servise) GetTimezoneByZone(zone int) (*PresentedTimezoneData, error) {
	whereCase := fmt.Sprintf("zone = %d", zone)
	return s.getRowWithWhereCase(whereCase)
}

// select a lot

func (s *servise) GetAllTimezones(limit_ofset ...int) (*[]PresentedTimezoneData, error) {
	return s.getRowsWithWhereCase("", limit_ofset...)
}

func (s *servise) GetSeveralTimezones(zones []int, limit_ofset ...int) (*[]PresentedTimezoneData, error) {
	zonesStr := make([]string, len(zones))
	for i, zone := range zones {
		zonesStr[i] = strconv.Itoa(zone)
	}
	whereCase := fmt.Sprintf("zone in (%s)", strings.Join(zonesStr, ", "))
	return s.getRowsWithWhereCase(whereCase, limit_ofset...)
}

// add

func (s *servise) checkTimezoneExist(zone int) bool {
	data, _ := s.GetTimezoneByZone(zone)
	return data != nil
}

func (s *servise) AddTimezone(data *MainTimezoneData) (*PresentedTimezoneData, error) {
	if s.checkTimezoneExist(data.Zone) {
		return nil, &timezoneExistError{}
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
	pdata := PresentedTimezoneDataFromMain(data, id)
	return &pdata, nil
}

// delete

func (*servise) DeleteTimezoneById(id int64) (bool, error) {
	stmt := fmt.Sprintf(`DELETE FROM timezone WHERE id = %d LIMIT 1`, id)
	db := database.Get_db()
	defer db.Close()

	_, err := db.Exec(stmt)
	if err != nil {
		return false, err
	}
	return true, err
}

func (*servise) DeleteTimezoneByZone(zone int) (bool, error) {
	stmt := fmt.Sprintf(`DELETE FROM timezone WHERE zone = %d`, zone)
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
