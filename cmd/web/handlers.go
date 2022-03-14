package main

import (
	"net/http"
	"strconv"
	"time"
)

const (
	Year = iota
	Month
)

func (app *application) howMuch(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		pd, ok := app.getPayDate(w, r)
		if !ok {
			return
		}
		data := app.getResponseData(pd, Month)
		err := app.writeJSON(w, http.StatusOK, data, nil)
		if err != nil {
			app.serverError(w, err)
		}
	default:
		w.Header().Set("Allow", "GET")
		app.clientError(w, 405)
	}

}

func (app *application) howMany(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		pd, ok := app.getPayDate(w, r)
		if !ok {
			return
		}
		data := app.getResponseData(pd, Year)
		err := app.writeJSON(w, http.StatusOK, data, nil)
		if err != nil {
			app.serverError(w, err)
		}
	default:
		w.Header().Set("Allow", "GET")
		app.clientError(w, 405)
	}
}

func (app *application) getResponseData(payDate int, period int) (data interface{}) {
	switch period {
	case Year:
		dates := whenSalaryYear(payDate)
		y := Yearly{
			Dates: dates,
		}
		return y
	case Month:
		days, date := whenSalaryMonth(payDate)
		m := Monthly{
			NrOfDays: days,
			Date:     date,
		}
		return m
	}
	return struct{}{}
}

// check if there are multiple queries as well - should it be invalidated?
func (app *application) getPayDate(w http.ResponseWriter, r *http.Request) (int, bool) {
	pd, err := strconv.Atoi(r.URL.Query().Get("pay-day"))
	y, m, _ := time.Now().Date()
	lastDayOfMonth := Date(y, m+1, 0).Day()

	if err != nil || pd < 1 || pd > lastDayOfMonth {
		app.clientError(w, 400)
		return 0, false
	}
	return pd, true
}
