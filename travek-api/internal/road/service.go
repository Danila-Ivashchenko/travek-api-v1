package road

import (
	"fmt"
	"travek-api/pkg/database"
)

type rowScaner interface {
	Scan(dest ...any) error
}

func presentedDataFromRowScaner(row rowScaner) (*PresentedRoadData, error) {
	data := &PresentedRoadData{}
	err := row.Scan(&data.Id, &data.FirstCountry, &data.SecondCountry, &data.Transport, &data.TimeHours)
	if err != nil {
		return nil, err
	}
	return data, nil
}

type Servise interface {
	AddRoad(data *MainRoadData) (*PresentedRoadData, error)
	AddRoads(data *[]MainRoadData) (int64, error)

	GetRoadById(id int64) (*PresentedRoadData, error)
	GetRoadByCountyId(id int64) (*PresentedRoadData, error)

	GetAllRoads(limit_ofset ...int) (*[]PresentedRoadData, error)
	GetAllRoadsByCountriesIds(first_id, second_id int64) (*[]PresentedRoadData, error)

	GetAllRoadsByCountryId(countryId int64) (*[]PresentedRoadData, error)
	GetAllRoadsByCountryName(name string) (*[]PresentedRoadData, error)

	DeleteRoadById(id int64) (bool, error)
	DeleteRoadByCountriesIds(first_id, second_id int64) (bool, error)
}

type servise struct{}

// primal selects

func (*servise) getRowWithWhereCase(whereCase string) (*PresentedRoadData, error) {
	stmt := "SELECT * FROM road "
	if whereCase != "" {
		stmt += fmt.Sprintf("WHERE %s", whereCase)
	}
	db := database.Get_db()
	defer db.Close()

	fmt.Println(stmt)
	row := db.QueryRow(stmt)
	data, err := presentedDataFromRowScaner(row)
	return data, err
}

func (*servise) getRowsWithWhereCase(whereCase string, limit_ofset ...int) (*[]PresentedRoadData, error) {
	stmt := "SELECT * FROM road "
	if whereCase != "" {
		stmt += fmt.Sprintf("WHERE %s", whereCase)
	}
	if len(limit_ofset) >= 1 {
		stmt += fmt.Sprintf(` LIMIT %d`, limit_ofset[0])
	}
	if len(limit_ofset) >= 2 {
		stmt += fmt.Sprintf(`, %d`, limit_ofset[1])
	}
	fmt.Println(stmt)
	db := database.Get_db()
	defer db.Close()

	rows, err := db.Query(stmt)
	if err != nil {
		return nil, err
	}
	pdatas := []PresentedRoadData{}
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

func (s *servise) GetRoadByData(data *MainRoadData) (*PresentedRoadData, error) {
	whereCase := fmt.Sprintf(`(first_country = %d AND second_country = %d) OR (first_country = %d AND second_country = %d) AND transport = "%s" AND time_hours = %d`, data.FirstCountry, data.SecondCountry, data.SecondCountry, data.FirstCountry, data.Transport, data.TimeHours)
	return s.getRowWithWhereCase(whereCase)
}

func (s *servise) GetRoadByCountyId(id int64) (*PresentedRoadData, error) {
	whereCase := fmt.Sprintf(`first_country = %d or second_country = %d`, id, id)
	return s.getRowWithWhereCase(whereCase)
}

func (s *servise) GetRoadById(id int64) (*PresentedRoadData, error) {
	whereCase := fmt.Sprintf(`id = %d`, id)
	return s.getRowWithWhereCase(whereCase)
}

func (s *servise) checkRoadExist(data *MainRoadData) bool {
	result, _ := s.GetRoadByData(data)
	return result != nil
}

// select a lot

func (s *servise) GetAllRoads(limit_ofset ...int) (*[]PresentedRoadData, error) {
	return s.getRowsWithWhereCase("", limit_ofset...)
}

func (s *servise) GetAllRoadsByCountryId(countryId int64) (*[]PresentedRoadData, error) {
	whereCase := fmt.Sprintf("first_country = %d OR second_country = %d", countryId, countryId)
	return s.getRowsWithWhereCase(whereCase)
}

func (s *servise) GetAllRoadsByCountriesIds(first_id, second_id int64) (*[]PresentedRoadData, error) {
	whereCase := fmt.Sprintf(`(first_country = %d AND second_country = %d) OR (first_country = %d AND second_country = %d)`, first_id, second_id, second_id, first_id)
	return s.getRowsWithWhereCase(whereCase)
}

func (s *servise) GetAllRoadsByCountryName(name string) (*[]PresentedRoadData, error) {
	whereCase := fmt.Sprintf(`SELECT road.id, road.first_country, road.second_country, road.transport, road.time_hourse FROM road, country WEHRE country.name = "%s" AND country.id IN (road.first_country, road.second_country)`, name)
	return s.getRowsWithWhereCase(whereCase)
}

// add

func (s *servise) AddRoad(data *MainRoadData) (*PresentedRoadData, error) {
	if s.checkRoadExist(data) {
		return nil, &roadExistError{}
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
	pdata := PresentedRoadDataFromMain(data, id)
	return &pdata, nil
}

func (s *servise) AddRoads(datas *[]MainRoadData) (int64, error) {
	for _, data := range *datas {
		if s.checkRoadExist(&data) {
			return -1, &roadExistError{}
		}
	}
	db := database.Get_db()
	defer db.Close()

	stmt := (*datas)[0].sqlInsertString()
	for i := 1; i < len(*datas); i++ {
		stmt += ", " + (*datas)[i].sqlValuesString()
	}
	result, err := db.Exec(stmt)
	if err != nil {
		return -1, err
	}
	return result.RowsAffected()
}

// delete

func (*servise) DeleteRoadById(id int64) (bool, error) {
	stmt := fmt.Sprintf(`DELETE FROM road WHERE id = %d LIMIT 1`, id)
	db := database.Get_db()
	defer db.Close()

	_, err := db.Exec(stmt)
	if err != nil {
		return false, err
	}
	return true, err
}

func (*servise) DeleteRoadByCountriesIds(first_id, second_id int64) (bool, error) {
	stmt := fmt.Sprintf(`DELETE FROM road WHERE (first_country = %d AND second_country = %d) OR (first_country = %d AND second_country = %d)`, first_id, second_id, first_id, second_id)
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
