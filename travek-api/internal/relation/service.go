package relation

import (
	"fmt"
	"travek-api/pkg/database"
)

type rowScaner interface {
	Scan(dest ...any) error
}

func presentedDataFromRowScaner(row rowScaner) (*PresentedRelationData, error) {
	data := &PresentedRelationData{}
	err := row.Scan(&data.Id, &data.FirstCountry, &data.SecondCountry, &data.FreeEntry, &data.PossibilityToStayForever)
	if err != nil {
		return nil, err
	}
	return data, nil
}

type Servise interface {
	AddRelation(data *MainRelationData) (*PresentedRelationData, error)

	GetAllRelations(limit_ofset ...int) (*[]PresentedRelationData, error)
	GetRelationById(id int64) (*PresentedRelationData, error)
	GetRelationByCountyId(id int64) (*PresentedRelationData, error)
	GetAllRelationsByCountryId(countryId int64) (*[]PresentedRelationData, error)
	GetAllRelationsByCountryName(name string) (*[]PresentedRelationData, error)

	DeleteRelationById(id int64) (bool, error)
	DeleteRelationByCountriesIds(first_id, second_id int64) (bool, error)
}

type servise struct{}

func (*servise) GetRelationByTwoIds(first_id, second_id int64) (*PresentedRelationData, error) {
	stmt := fmt.Sprintf(`SELECT * FROM relation WHERE (first_country = %d AND second_country = %d) OR (first_country = %d AND second_country = %d)`, first_id, second_id, first_id, second_id)
	db := database.Get_db()
	defer db.Close()

	row := db.QueryRow(stmt)
	data, err := presentedDataFromRowScaner(row)
	return data, err
}

func (*servise) GetRelationByCountyId(id int64) (*PresentedRelationData, error) {
	stmt := fmt.Sprintf(`SELECT * FROM relation WHERE first_country = %d or second_country = %d`, id, id)
	db := database.Get_db()
	defer db.Close()

	row := db.QueryRow(stmt)
	data, err := presentedDataFromRowScaner(row)
	return data, err
}

func (*servise) GetRelationById(id int64) (*PresentedRelationData, error) {
	stmt := fmt.Sprintf(`SELECT * FROM relation WHERE id = %d`, id)
	db := database.Get_db()
	defer db.Close()

	row := db.QueryRow(stmt)
	data, err := presentedDataFromRowScaner(row)
	return data, err
}

func (s *servise) checkRelationExist(first_id, second_id int64) bool {
	data, _ := s.GetRelationByTwoIds(first_id, second_id)
	return data != nil
}

func (s *servise) AddRelation(data *MainRelationData) (*PresentedRelationData, error) {
	if s.checkRelationExist(data.FirstCountry, data.SecondCountry) {
		return nil, &relationExistError{}
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
	pdata := PresentedRelationDataFromMain(data, id)
	return &pdata, nil
}

func (*servise) GetAllRelations(limit_ofset ...int) (*[]PresentedRelationData, error) {
	stmt := "SELECT * FROM relation"
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
	pdatas := []PresentedRelationData{}
	for rows.Next() {
		pdata, err := presentedDataFromRowScaner(rows)
		if err != nil {
			break
		}
		pdatas = append(pdatas, *pdata)
	}
	return &pdatas, nil
}

func (*servise) GetAllRelationsByCountryId(countryId int64) (*[]PresentedRelationData, error) {
	stmt := fmt.Sprintf("SELECT * FROM relation WHERE first_country = %d OR second_country = %d", countryId, countryId)
	db := database.Get_db()
	defer db.Close()

	rows, err := db.Query(stmt)
	if err != nil {
		return nil, err
	}
	pdatas := []PresentedRelationData{}
	for rows.Next() {
		pdata, err := presentedDataFromRowScaner(rows)
		if err != nil {
			break
		}
		pdatas = append(pdatas, *pdata)
	}
	return &pdatas, nil
}
func (*servise) GetAllRelationsByCountryName(name string) (*[]PresentedRelationData, error) {
	stmt := fmt.Sprintf(`SELECT relation.id, relation.first_country, relation.second_country, relation.free_entry, relation.possibility_to_stay_forever FROM relation, country WEHRE country.name = "%s" AND country.id IN (relation.first_country, relation.second_country)`, name)
	db := database.Get_db()
	defer db.Close()

	rows, err := db.Query(stmt)
	if err != nil {
		return nil, err
	}
	pdatas := []PresentedRelationData{}
	for rows.Next() {
		pdata, err := presentedDataFromRowScaner(rows)
		if err != nil {
			break
		}
		pdatas = append(pdatas, *pdata)
	}
	return &pdatas, nil
}

func (*servise) DeleteRelationById(id int64) (bool, error) {
	stmt := fmt.Sprintf(`DELETE FROM relation WHERE id = %d LIMIT 1`, id)
	db := database.Get_db()
	defer db.Close()

	_, err := db.Exec(stmt)
	if err != nil {
		return false, err
	}
	return true, err
}

func (*servise) DeleteRelationByCountriesIds(first_id, second_id int64) (bool, error) {
	stmt := fmt.Sprintf(`DELETE FROM relation WHERE (first_country = %d AND second_country = %d) OR (first_country = %d AND second_country = %d)`, first_id, second_id, first_id, second_id)
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
