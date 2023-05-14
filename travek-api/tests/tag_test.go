package tests

import (
	"testing"
	"travek-api/internal/tags"
)

func TestTagInsertDelete(t *testing.T) {
	ts := tags.GetService()
	timezones, err := ts.GetAllTags()
	if err != nil {
		t.Errorf("Error: %s", err.Error())
	}
	timezonesLen := len(*timezones)
	pdatas := []tags.PresentedTagData{}

	tagsToAdd := []tags.MainTagData{
		tags.MainTagData{
			Name: "a",
		},
		tags.MainTagData{
			Name: "b",
		},
		tags.MainTagData{
			Name: "c",
		},
		tags.MainTagData{
			Name: "d",
		},
	}
	for _, data := range tagsToAdd {
		pdata, err := ts.AddTag(&data)
		if err != nil {
			t.Errorf("Error: %s", err.Error())
		}
		pdatas = append(pdatas, *pdata)
	}

	for i, pdata := range pdatas {
		var (
			flag bool
			err  error
		)
		if i%2 == 0 {
			flag, err = ts.DeleteTagById(pdata.Id)
		} else {
			flag, err = ts.DeleteTagByName(pdata.Name)
		}
		if !flag || err != nil {
			t.Errorf("Error: %s", err.Error())
		}
	}
	newTimezonesLen, err := ts.GetAllTags()
	if err != nil {
		t.Errorf("Error: %s", err.Error())
	}
	if len(*newTimezonesLen) != timezonesLen {
		t.Errorf("Error: %s", "Bad get all")
	}
	//fmt.Println(ts.GetSeveralTimezones([]int{-1, 0, 1, 2}))
}
