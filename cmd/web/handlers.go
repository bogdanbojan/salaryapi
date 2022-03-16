package main

import (
	"net/http"
	"strconv"
	"time"
)

// howMuch is the handler for the GET request asking the next salary date and how many days
// you have until then. It only works with GET requests.
func (app *application) howMuch(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		pd, ok := app.getPayDay(w, r)
		if !ok {
			app.clientError(w, 400)
			return
		}
		data := app.getResponseData(pd, Month)
		err := app.writeJSON(w, http.StatusOK, data, nil)
		if err != nil {
			app.serverError(w, err)
		}
	default:
		w.Header().Set("Allow", "GET")
		app.clientError(w, http.StatusMethodNotAllowed)
	}

}

// howMuch is the handler for the GET request asking the next salary dates in the
// current year. It only works with GET requests.
func (app *application) howMany(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		pd, ok := app.getPayDay(w, r)
		if !ok {
			app.clientError(w, 400)
			return
		}
		data := app.getResponseData(pd, Year)
		err := app.writeJSON(w, http.StatusOK, data, nil)
		if err != nil {
			app.serverError(w, err)
		}
	default:
		w.Header().Set("Allow", "GET")
		app.clientError(w, http.StatusMethodNotAllowed)
	}
}

// getResponseData gets the type of data that needs to be written to the JSON response.
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

// getPayDay extracts the pay day from the query. It checks if it is valid and returns a
// tuple with the pay day and a flag which tells you if the input was valid or not.
func (app *application) getPayDay(w http.ResponseWriter, r *http.Request) (int, bool) {
	pd, err := strconv.Atoi(r.URL.Query().Get("pay-day"))
	y, m, _ := time.Now().Date()
	lastDayOfMonth := Date(y, m+1, 0).Day()

	if err != nil || pd < 1 || pd > lastDayOfMonth {
		return 0, false
	}
	return pd, true
}
