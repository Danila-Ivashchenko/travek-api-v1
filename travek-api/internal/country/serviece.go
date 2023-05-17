package country

import (
	"fmt"
	"travek-api/pkg/database"
)

type rowScaner interface {
	Scan(dest ...any) error
}

func presentedDataFromRowScaner(row rowScaner) (*PresentedCountryData, error) {
	data := &PresentedCountryData{}
	err := row.Scan(&data.Id, &data.Name, &data.Language)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func allDataFromRowScaner(row rowScaner) (*AllCountryData, error) {
	data := &AllCountryData{}
	err := row.Scan(&data.Id, &data.Name, &data.Language, &data.Description)
	if err != nil {
		return nil, err
	}
	return data, nil
}

type Servise interface {
	AddCountry(data *MainCountryData) (*PresentedCountryData, error)
	AddCountries(data *[]MainCountryData) (int64, error)

	GetCountyByName(name string) (*PresentedCountryData, error)
	GetCountyById(id int64) (*PresentedCountryData, error)
	GetAllCountryDataById(id int64) (*AllCountryData, error)
	GetAllCountryDataByName(name string) (*AllCountryData, error)

	GetCountriesWithWhere(whereCase string, limit_ofset ...int) (*[]PresentedCountryData, error)
	GetAllCountries(limit_ofset ...int) (*[]PresentedCountryData, error)
	GetCountriesByLanguage(language string, limit_ofset ...int) (*[]PresentedCountryData, error)

	GetCountriesWithHaving(havingCase string, limit_ofset ...int) (*[]PresentedCountryData, error)

	DeleteCountryById(id int64) (bool, error)
	DeleteCountryByName(name string) (bool, error)
}

type servise struct{}

// primal select

func (*servise) getRowWithWhereCase(whereCase string) (*PresentedCountryData, error) {
	stmt := "SELECT country.id, country.name, country.language FROM country"
	if whereCase != "" {
		stmt += fmt.Sprintf(" WHERE %s", whereCase)
	}
	db := database.Get_db()
	defer db.Close()

	row := db.QueryRow(stmt)
	data, err := presentedDataFromRowScaner(row)
	return data, err
}

func (*servise) getAllRowWithWhereCase(whereCase string) (*AllCountryData, error) {
	stmt := "SELECT * FROM country"
	if whereCase != "" {
		stmt += fmt.Sprintf(" WHERE %s", whereCase)
	}
	db := database.Get_db()
	defer db.Close()

	row := db.QueryRow(stmt)
	data, err := allDataFromRowScaner(row)
	return data, err
}

func (*servise) getRowsWithWhereCase(whereCase string, limit_ofset ...int) (*[]PresentedCountryData, error) {
	stmt := "SELECT country.id, country.name, country.language FROM country"
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
	counties := []PresentedCountryData{}
	for rows.Next() {
		data, err := presentedDataFromRowScaner(rows)
		if err != nil {
			break
		}
		counties = append(counties, *data)
	}
	return &counties, nil
}

func (*servise) GetCountriesWithHaving(havingCase string, limit_ofset ...int) (*[]PresentedCountryData, error) {
	stmt := "SELECT country.id, country.name, country.language FROM country"
	if havingCase != "" {
		stmt += fmt.Sprintf(" HAVING %s", havingCase)
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
	counties := []PresentedCountryData{}
	for rows.Next() {
		data, err := presentedDataFromRowScaner(rows)
		if err != nil {
			break
		}
		counties = append(counties, *data)
	}
	return &counties, nil
}

// select one

func (s *servise) GetAllCountryDataById(id int64) (*AllCountryData, error) {
	whereCase := fmt.Sprintf("id = %d", id)
	return s.getAllRowWithWhereCase(whereCase)
}
func (s *servise) GetAllCountryDataByName(name string) (*AllCountryData, error) {
	whereCase := fmt.Sprintf(`name = "%s"`, name)
	return s.getAllRowWithWhereCase(whereCase)
}

func (s *servise) GetCountyByName(name string) (*PresentedCountryData, error) {
	whereCase := fmt.Sprintf(`name = "%s"`, name)
	return s.getRowWithWhereCase(whereCase)
}

func (s *servise) GetCountyById(id int64) (*PresentedCountryData, error) {
	whereCase := fmt.Sprintf("id = %d", id)
	return s.getRowWithWhereCase(whereCase)
}

// select a lot

func (s *servise) GetAllCountries(limit_ofset ...int) (*[]PresentedCountryData, error) {
	return s.getRowsWithWhereCase("", limit_ofset...)
}

func (s *servise) GetCountriesByLanguage(language string, limit_ofset ...int) (*[]PresentedCountryData, error) {
	whereCase := fmt.Sprintf("language = %s", language)
	return s.getRowsWithWhereCase(whereCase, limit_ofset...)
}

func (s *servise) checkCountryExist(name string) bool {
	data, _ := s.GetCountyByName(name)
	return data != nil
}

func (s *servise) GetCountriesWithWhere(whereCase string, limit_ofset ...int) (*[]PresentedCountryData, error) {
	return s.getRowsWithWhereCase(whereCase, limit_ofset...)
}

// add

func (s *servise) AddCountry(data *MainCountryData) (*PresentedCountryData, error) {
	if s.checkCountryExist(data.Name) {
		return nil, &countryExistError{}
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
	pdata := PresentedCountryDataFromMain(data, id)
	return &pdata, nil
}

func (s *servise) AddCountries(datas *[]MainCountryData) (int64, error) {
	for _, data := range *datas {
		if s.checkCountryExist(data.Name) {
			return 0, &countryExistError{}
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
		return 0, err
	}
	count, err := result.RowsAffected()
	return count, err
}

// delete

func (*servise) DeleteCountryById(id int64) (bool, error) {
	stmt := fmt.Sprintf(`DELETE FROM country WHERE id = %d LIMIT 1`, id)
	db := database.Get_db()
	defer db.Close()

	_, err := db.Exec(stmt)
	if err != nil {
		return false, err
	}
	return true, err
}

func (*servise) DeleteCountryByName(name string) (bool, error) {
	stmt := fmt.Sprintf(`DELETE FROM country WHERE name = "%s" LIMIT 1`, name)
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
