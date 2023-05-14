package agragator

import (
	"fmt"
	"strings"
	"travek-api/internal/country"
	"travek-api/internal/relation"
	"travek-api/internal/road"
	"travek-api/internal/tags"
	"travek-api/internal/timezone"
)

// country

type BigCountryAddData struct {
	country.MainCountryData
	TimeZones []int    `json:"time_zones"`
	Tags      []string `json:"tags"`
}

func (bd *BigCountryAddData) ExtractCountryData() country.MainCountryData {
	return bd.MainCountryData
}

func (bd *BigCountryAddData) Validate() bool {
	return bd.Name != "" && bd.Language != "" && bd.Continent != "" && len(bd.TimeZones) > 0
}

func (bd *BigCountryAddData) ExtractTimezones() *[]timezone.PresentedTimezoneData {
	ts := timezone.GetService()
	data, _ := ts.GetSeveralTimezones(bd.TimeZones)
	return data
}

func (bd *BigCountryAddData) ExtractTags() *[]tags.PresentedTagData {
	ts := tags.GetService()
	data, _ := ts.GetSeveralTags(bd.Tags)
	return data
}

// relation

type BigRelationAddData struct {
	FirstCountry             string `json:"first_country"`
	SecondCountry            string `json:"second_country"`
	FreeEntry                bool   `json:"free_entry"`
	PossibilityToStayForever bool   `json:"possibility_to_stay_forever"`
}

func (bd *BigRelationAddData) ExtractMainRelationData() (*relation.MainRelationData, error) {
	cs := country.GetService()
	mdata := relation.MainRelationData{}
	cdata, err := cs.GetCountyByName(bd.FirstCountry)
	if err != nil {
		return nil, err
	}
	mdata.FirstCountry = cdata.Id
	cdata, err = cs.GetCountyByName(bd.SecondCountry)
	if err != nil {
		return nil, err
	}
	mdata.SecondCountry = cdata.Id
	mdata.FreeEntry = bd.FreeEntry
	mdata.PossibilityToStayForever = bd.PossibilityToStayForever
	return &mdata, nil
}

// road

type BigRoadsAddData struct {
	FirstCountry  string         `json:"first_country"`
	SecondCountry string         `json:"second_country"`
	Roads         map[string]int `json:"roads"`
}

func (bd *BigRoadsAddData) ExtractMainRoadsData() (*[]road.MainRoadData, error) {
	datas := []road.MainRoadData{}
	cs := country.GetService()
	for key, value := range bd.Roads {
		data := road.MainRoadData{}
		cdata, err := cs.GetCountyByName(bd.FirstCountry)
		if err != nil {
			return nil, err
		}
		data.FirstCountry = cdata.Id
		cdata, err = cs.GetCountyByName(bd.SecondCountry)
		if err != nil {
			return nil, err
		}
		data.SecondCountry = cdata.Id
		data.Transport = key
		data.TimeHours = value
		datas = append(datas, data)
	}
	return &datas, nil
}

// tags

type TagsToCountryData struct {
	Country string   `json:"country"`
	Tags    []string `json:"tags"`
}

func (bd *TagsToCountryData) ExtractCountryData() (*country.PresentedCountryData, error) {
	pdata, err := country.GetService().GetCountyByName(bd.Country)
	return pdata, err
}

func (bd *TagsToCountryData) ExtractTags() *[]tags.PresentedTagData {
	ts := tags.GetService()
	data, _ := ts.GetSeveralTags(bd.Tags)
	return data
}

// criterias

type CriteriasData struct {
	SourseCountry            int64    `json:"sourse_country"`
	Languages                []string `json:"languages"`
	Transports               []string `json:"transports"`
	MaxTime                  int      `json:"max_time"`
	TimeZones                []int    `json:"time_zones"`
	FreeEntry                bool     `json:"free_entry"`
	PossibilityToStayForever bool     `json:"possibility_to_stay_forever"`
}

func (c *CriteriasData) prepareTransport() []string {
	transports := c.Transports
	for i, data := range transports {
		transports[i] = fmt.Sprintf("'%s'", data)
	}
	return transports
}

func (c *CriteriasData) extractRelationsSQL() string {
	stmt := fmt.Sprintf(`SELECT CASE WHEN relation.first_country = %d THEN relation.second_country WHEN relation.second_country = %d THEN relation.first_country END AS country_id FROM relation WHERE`, c.SourseCountry, c.SourseCountry)
	stmt += fmt.Sprintf(` (relation.first_country = %d OR relation.second_country = %d)`, c.SourseCountry, c.SourseCountry)
	if c.FreeEntry {
		stmt += fmt.Sprintf(` AND (relation.free_entry = %t)`, c.FreeEntry)
	}
	if c.PossibilityToStayForever {
		stmt += fmt.Sprintf(` AND (relation.possibility_to_stay_forever = %t)`, c.PossibilityToStayForever)
	}
	return stmt
}

func (c *CriteriasData) extractRoadsSQL() string {
	stmt := fmt.Sprintf(`SELECT CASE WHEN road.first_country = %d THEN road.second_country WHEN road.second_country = %d THEN road.first_country END AS country_id FROM road WHERE`, c.SourseCountry, c.SourseCountry)
	stmt += fmt.Sprintf(` (road.first_country = %d OR road.second_country = %d)`, c.SourseCountry, c.SourseCountry)
	if len(c.Transports) > 0 {
		stmt += fmt.Sprintf(` AND (road.transport IN(%s))`, strings.Join(c.prepareTransport(), ", "))
	}
	if c.MaxTime > 0 {
		stmt += fmt.Sprintf(` AND (road.time_hours <= %d)`, c.MaxTime)
	}
	return stmt
}

func (c *CriteriasData) addHaving(stmt, havingCase string) string {
	return stmt + fmt.Sprintf(` HAVING country_id IN(%s)`, havingCase)
}

func (c *CriteriasData) ExtractSQL() string {
	stmt := c.extractRelationsSQL()
	return c.addHaving(stmt, c.extractRoadsSQL())
}

func (c *CriteriasData) ExtractHavingCase(havingCase string) string {
	return havingCase + c.ExtractSQL()
}
