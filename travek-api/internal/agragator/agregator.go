package agragator

import (
	"travek-api/internal/country"
	"travek-api/internal/country_tags"
	"travek-api/internal/country_timezone"
	"travek-api/internal/relation"
	"travek-api/internal/road"
	"travek-api/internal/tags"
	"travek-api/internal/timezone"
)

type Agragator struct {
	Cs    country.Servise
	Rels  relation.Servise
	Roads road.Servise
	Ts    timezone.Servise
	Cts   country_timezone.Servise
	Tags  tags.Servise
	Ctags country_tags.Servise
}

// add funcs

func (a *Agragator) AddCountry(data *BigCountryAddData) (*country.PresentedCountryData, error) {
	cdata := data.ExtractCountryData()
	timeData := data.ExtractTimezones()

	pcdata, err := a.Cs.AddCountry(&cdata)
	if err != nil {
		return nil, err
	}
	_, err = a.Cts.AddCountryTimezones(pcdata.Id, timeData)
	if err != nil {
		return pcdata, err
	}
	if len(data.Tags) > 0 {
		tagsData := data.ExtractTags()
		_, err = a.Ctags.AddCountyTags(pcdata.Id, tagsData)
	}
	return pcdata, err
}

func (a *Agragator) AddRelations(data *BigRelationAddData) (*relation.PresentedRelationData, error) {
	rdata, err := data.ExtractMainRelationData()
	if err != nil {
		return nil, err
	}

	pdata, err := a.Rels.AddRelation(rdata)
	return pdata, err
}

func (a *Agragator) AddRoads(data *BigRoadsAddData) (*[]road.PresentedRoadData, error) {
	rdatas, err := data.ExtractMainRoadsData()
	if err != nil {
		return nil, err
	}
	_, err = a.Roads.AddRoads(rdatas)
	if err != nil {
		return nil, err
	}
	first_country := (*rdatas)[0].FirstCountry
	second_country := (*rdatas)[0].SecondCountry
	pdats, err := a.Roads.GetAllRoadsByCountriesIds(first_country, second_country)
	return pdats, err
}

func (a *Agragator) AddTags(data *[]tags.MainTagData) (*[]tags.PresentedTagData, error) {
	_, err := a.Tags.AddTags(data)
	if err != nil {
		return nil, err
	}
	tagsStr := make([]string, len(*data))
	for i := range *data {
		tagsStr[i] = (*data)[i].Name
	}
	return a.Tags.GetSeveralTags(tagsStr)
}

func (a *Agragator) AddTagsToCountry(data *TagsToCountryData) (*[]country_tags.PresentedCountyTagsData, error) {
	tags := data.ExtractTags()
	countryData, err := data.ExtractCountryData()

	if err != nil {
		return nil, err
	}
	_, err = a.Ctags.AddCountyTags(countryData.Id, tags)
	if err != nil {
		return nil, err
	}

	return a.Ctags.GetCountyTagsByCountryId(countryData.Id)
}

// get results

func (a *Agragator) GetSuitable(data *CriteriasData) (*[]country.PresentedCountryData, error) {
	// havingCase := data.ExtractSQL()
	// stmt := ""
	// if havingCase != "" {
	// 	stmt = fmt.Sprintf(`country.id IN(%s)`, havingCase)
	// }
	// return a.Cs.GetCountriesWithHaving(stmt)
	whereCase := data.getWhereCase()

	return a.Cs.GetCountriesWithWhere(whereCase)
}

func GetAgregator() Agragator {
	a := Agragator{}
	a.Cs = country.GetService()
	a.Cts = country_timezone.GetService()
	a.Rels = relation.GetService()
	a.Roads = road.GetService()
	a.Ts = timezone.GetService()
	a.Tags = tags.GetService()
	a.Ctags = country_tags.GetService()

	return a
}
