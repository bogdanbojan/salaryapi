package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
)

func TestWriteJSON(t *testing.T) {
	for _, mc := range writeJSONCases.monthlyCases {
		rr := httptest.NewRecorder()
		err := app.writeJSON(rr, http.StatusOK, mc, nil)
		if err != nil {
			return
		}
		rs := rr.Result()
		body, err := ioutil.ReadAll(rs.Body)
		if err != nil {
			return
		}
		bodyCase, _ := json.MarshalIndent(mc, "", "\t")
		bodyCase = append(bodyCase, '\n')

		assertWriteJSON(t, body, bodyCase)
		err = rs.Body.Close()
		if err != nil {
			return
		}
	}
	for _, yc := range writeJSONCases.yearlyCases {
		rr := httptest.NewRecorder()
		err := app.writeJSON(rr, http.StatusOK, yc, nil)
		if err != nil {
			return
		}
		rs := rr.Result()
		body, err := ioutil.ReadAll(rs.Body)
		if err != nil {
			return
		}
		bodyCase, _ := json.MarshalIndent(yc, "", "\t")
		bodyCase = append(bodyCase, '\n')

		assertWriteJSON(t, body, bodyCase)
		err = rs.Body.Close()
		if err != nil {
			return
		}
	}

}

func assertWriteJSON(t testing.TB, got, want []byte) {
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %s, want %s", got, want)
	}
}

var writeJSONCases = struct {
	monthlyCases []Monthly
	yearlyCases  []Yearly
	//wantMonthlyCases []string
	//wantYearlyCases  []string
}{
	monthlyCases: []Monthly{
		{
			NrOfDays: 16,
			Date:     "Tuesday, 15-Mar-22",
		},
	},
	yearlyCases: []Yearly{
		{
			Dates: []string{
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
		},
	},
}
