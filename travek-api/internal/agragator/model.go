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
	return bd.Name != "" && bd.Language != "" && bd.Description != "" && len(bd.TimeZones) > 0
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
	InternalPassport         bool   `json:"internal_passport"`
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
	mdata.PossibilityToStayForever = bd.PossibilityToStayForever
	mdata.InternalPassport = bd.InternalPassport
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
	Tags                     []int    `json:"tags"`
	InternalPassport         bool     `json:"internal_passport"`
	PossibilityToStayForever bool     `json:"possibility_to_stay_forever"`
}

func (c *CriteriasData) extractTagsCase() string {
	if len(c.Tags) == 0 {
		return ""
	}
	tags := ""
	for i, tag := range c.Tags {
		tagStr := fmt.Sprintf(`%d IN((select tag_id from country_tags where country_id = c_id))`, tag)
		if i+1 != len(c.Tags) {
			tags += fmt.Sprintf("%s AND ", tagStr)
		} else {
			tags += tagStr
		}
	}

	stmt := fmt.Sprintf(`(SELECT country_id as c_id FROM country_tags GROUP BY c_id HAVING %s)`, tags)
	return fmt.Sprintf("country.id IN(%s)", stmt)
}

func (c *CriteriasData) prepareTransport() []string {
	transports := c.Transports
	for i, data := range transports {
		transports[i] = fmt.Sprintf("'%s'", data)
	}
	return transports
}

func (c *CriteriasData) extractRelationsCase() string {
	// if !c.InternalPassport && !c.PossibilityToStayForever {
	// 	return ""
	// }
	stmt := fmt.Sprintf(`SELECT CASE WHEN relation.first_country = %d THEN relation.second_country WHEN relation.second_country = %d THEN relation.first_country END AS country_id FROM relation WHERE`, c.SourseCountry, c.SourseCountry)
	stmt += fmt.Sprintf(` (relation.first_country = %d OR relation.second_country = %d)`, c.SourseCountry, c.SourseCountry)
	// if !c.InternalPassport {
	// 	stmt += fmt.Sprintf(` AND (relation.internal_passport = %t)`, c.InternalPassport)
	// }
	stmt += fmt.Sprintf(` AND (relation.internal_passport = %t)`, c.InternalPassport)
	if c.PossibilityToStayForever {
		stmt += fmt.Sprintf(` AND (relation.possibility_to_stay_forever = %t)`, c.PossibilityToStayForever)
	}

	return fmt.Sprintf("country.id IN(%s)", stmt)
}

func (c *CriteriasData) extractRoadsCase() string {
	if len(c.Transports) == 0 && c.MaxTime == 0 {
		return ""
	}
	stmt := fmt.Sprintf(`SELECT CASE WHEN road.first_country = %d THEN road.second_country WHEN road.second_country = %d THEN road.first_country END AS country_id FROM road WHERE`, c.SourseCountry, c.SourseCountry)
	stmt += fmt.Sprintf(` (road.first_country = %d OR road.second_country = %d)`, c.SourseCountry, c.SourseCountry)
	if len(c.Transports) > 0 {
		stmt += fmt.Sprintf(` AND (road.transport IN(%s))`, strings.Join(c.prepareTransport(), ", "))
	}
	if c.MaxTime > 0 {
		stmt += fmt.Sprintf(` AND (road.time_hours <= %d)`, c.MaxTime)
	}

	fmt.Println("road case", stmt)
	return fmt.Sprintf("country.id IN(%s)", stmt)
}

func (c *CriteriasData) getWhereCase() string {
	tagsCase := c.extractTagsCase()
	relCase := c.extractRelationsCase()
	roadCase := c.extractRoadsCase()
	stmt := ""
	if tagsCase != "" {
		stmt += tagsCase
	}
	if relCase != "" {
		if stmt != "" {
			stmt += " AND"
		}
		stmt += " (" + relCase
		stmt += fmt.Sprintf(` OR country.id = %d)`, c.SourseCountry)
	}
	if roadCase != "" {
		if stmt != "" {
			stmt += " AND"
		}
		stmt += " (" + roadCase
		stmt += fmt.Sprintf(` OR country.id = %d)`, c.SourseCountry)
	}
	// fmt.Println(stmt)
	return stmt
}

func (c *CriteriasData) addHaving(stmt, havingCase string) string {
	if havingCase != "" {
		return stmt + fmt.Sprintf(` HAVING country_id IN(%s)`, havingCase)
	} else {
		return stmt
	}
}

func (c *CriteriasData) ExtractSQL() string {
	stmt := c.extractRelationsCase()
	return c.addHaving(stmt, c.extractRoadsCase())
}

func (c *CriteriasData) ExtractHavingCase(havingCase string) string {
	return havingCase + c.ExtractSQL()
}
