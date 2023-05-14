package tags

import (
	"fmt"
	"strings"
	"travek-api/pkg/database"
)

type rowScaner interface {
	Scan(dest ...any) error
}

func presentedDataFromRowScaner(row rowScaner) (*PresentedTagData, error) {
	data := &PresentedTagData{}
	err := row.Scan(&data.Id, &data.Name)
	if err != nil {
		return nil, err
	}
	return data, nil
}

type Servise interface {
	AddTag(data *MainTagData) (*PresentedTagData, error)
	AddTags(datas *[]MainTagData) (int64, error)

	GetSeveralTags(tags []string, limit_ofset ...int) (*[]PresentedTagData, error)
	GetAllTags(limit_ofset ...int) (*[]PresentedTagData, error)

	GetTagById(id int64) (*PresentedTagData, error)
	GetTagByName(name string) (*PresentedTagData, error)

	DeleteTagById(id int64) (bool, error)
	DeleteTagByName(name string) (bool, error)
}

type servise struct{}

// primal select

func (*servise) getRowWithWhereCase(whereCase string) (*PresentedTagData, error) {
	stmt := "SELECT * FROM tags"
	if whereCase != "" {
		stmt += fmt.Sprintf(" WHERE %s", whereCase)
	}
	db := database.Get_db()
	defer db.Close()

	row := db.QueryRow(stmt)
	data, err := presentedDataFromRowScaner(row)
	return data, err
}

func (*servise) getRowsWithWhereCase(whereCase string, limit_ofset ...int) (*[]PresentedTagData, error) {
	stmt := "SELECT * FROM tags"
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
	pdatas := []PresentedTagData{}
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

func (s *servise) GetTagById(id int64) (*PresentedTagData, error) {
	whereCase := fmt.Sprintf("id = %d", id)
	return s.getRowWithWhereCase(whereCase)
}

func (s *servise) GetTagByName(name string) (*PresentedTagData, error) {
	whereCase := fmt.Sprintf(`name = "%s"`, name)
	return s.getRowWithWhereCase(whereCase)
}

// select a lot

func (s *servise) GetAllTags(limit_ofset ...int) (*[]PresentedTagData, error) {
	return s.getRowsWithWhereCase("", limit_ofset...)
}

func (s *servise) GetSeveralTags(tags []string, limit_ofset ...int) (*[]PresentedTagData, error) {
	tagsStr := make([]string, len(tags))
	for i, tag := range tags {
		tagsStr[i] = fmt.Sprintf(`"%s"`, tag)
	}
	whereCase := fmt.Sprintf("name in (%s)", strings.Join(tagsStr, ", "))
	return s.getRowsWithWhereCase(whereCase, limit_ofset...)
}

// add

func (s *servise) checkTagExist(name string) bool {
	data, _ := s.GetTagByName(name)
	return data != nil
}

func (s *servise) AddTag(data *MainTagData) (*PresentedTagData, error) {
	if s.checkTagExist(data.Name) {
		return nil, &tagExistError{}
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
	pdata := PresentedTagDataFromMain(data, id)
	return &pdata, nil
}

func (s *servise) AddTags(datas *[]MainTagData) (int64, error) {
	for _, data := range *datas {
		if s.checkTagExist(data.Name) {
			return -1, &tagExistError{}
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
	count, err := result.RowsAffected()
	if err != nil {
		return -1, err
	}

	return count, nil
}

// delete

func (*servise) DeleteTagById(id int64) (bool, error) {
	stmt := fmt.Sprintf(`DELETE FROM tags WHERE id = %d LIMIT 1`, id)
	db := database.Get_db()
	defer db.Close()

	_, err := db.Exec(stmt)
	if err != nil {
		return false, err
	}
	return true, err
}

func (*servise) DeleteTagByName(name string) (bool, error) {
	stmt := fmt.Sprintf(`DELETE FROM tags WHERE name = "%s"`, name)
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
