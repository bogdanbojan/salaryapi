package main

import (
	"reflect"
	"testing"
	"time"
)

// Current date for the tests.
const y, m, d = 2022, time.Month(03), 03

func TestUpToDays(t *testing.T) {
	currentDate := Date(y, m, d)
	salaryDate := Date(y, m, 20)

	daysCount := upToDays(currentDate, salaryDate)
	if daysCount != 17 {
		t.Errorf("got %d, want %d", daysCount, 17)
	}
}

func TestCheckWeekday(t *testing.T) {
	for i, wc := range checkWeekdayCases.wcase {
		date := checkWeekday(wc)
		assertDates(t, date, checkWeekdayCases.want[i])
	}
}

func assertDates(t testing.TB, date1, date2 time.Time) {
	t.Helper()
	if date1 != date2 {
		t.Errorf("got %s, want %s", date1, date2)
	}
}

var checkWeekdayCases = struct {
	wcase []time.Time
	want  []time.Time
}{
	wcase: []time.Time{
		Date(y, m, 19),
		Date(y, m, 10),
		Date(y, m, 26),
	},
	want: []time.Time{
		Date(y, m, 25),
		Date(y, m, 10),
		Date(y, m+1, 1),
	},
}

func TestYearlySalaryDates(t *testing.T) {
	ysd := getYearlySalaryDates(Date(y, m, 15))
	assertYearlySalaryDates(t, ysd, yearlySalaryDatesCases.want)
}

func assertYearlySalaryDates(t testing.TB, date1, date2 []string) {
	t.Helper()
	if !reflect.DeepEqual(date1, date2) {
		t.Errorf("got %s, want %s", date1, date2)

	}
}

var yearlySalaryDatesCases = struct {
	want []string
}{
	want: []string{
		"Tuesday, 15-Mar-22",
		"Friday, 15-Apr-22",
		"Friday, 20-May-22",
		"Wednesday, 15-Jun-22",
		"Friday, 15-Jul-22",
		"Monday, 15-Aug-22",
		"Thursday, 15-Sep-22",
		"Friday, 21-Oct-22",
		"Tuesday, 15-Nov-22",
		"Thursday, 15-Dec-22",
	},
}
