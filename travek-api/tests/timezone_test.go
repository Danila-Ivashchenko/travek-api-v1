package tests

import (
	"fmt"
	"testing"
	"travek-api/internal/timezone"
)

func TestTimezoneInsertDelete(t *testing.T) {
	ts := timezone.GetService()
	timezones, err := ts.GetAllTimezones()
	if err != nil {
		t.Errorf("Error: %s", err.Error())
	}
	timezonesLen := len(*timezones)
	pdatas := []timezone.PresentedTimezoneData{}
	for i := -12; i <= 14; i++ {
		mdata := timezone.MainTimezoneData{Zone: i}
		pdata, err := ts.AddTimezone(&mdata)
		if err != nil {
			t.Errorf("Error: %s", err.Error())
		}
		pdatas = append(pdatas, *pdata)
	}
	if len(pdatas) != 25 {
		t.Errorf("Error: expected 25, got %d", len(pdatas))
	}
	// for i, pdata := range pdatas {
	// 	var (
	// 		flag bool
	// 		err  error
	// 	)
	// 	if i%2 == 0 {
	// 		flag, err = ts.DeleteTimezoneById(pdata.Id)
	// 	} else {
	// 		flag, err = ts.DeleteTimezoneByZone(pdata.Zone)
	// 	}
	// 	if !flag || err != nil {
	// 		t.Errorf("Error: %s", err.Error())
	// 	}
	// }
	newTimezonesLen, err := ts.GetAllTimezones()
	if err != nil {
		t.Errorf("Error: %s", err.Error())
	}
	if len(*newTimezonesLen) != timezonesLen {
		t.Errorf("Error: %s", "Bad get all")
	}
	fmt.Println(ts.GetSeveralTimezones([]int{-1, 0, 1, 2}))
}
