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

func whenSalaryMonth(payDay int) (int, string) {
	currentDate, salaryDate := getAPIDates(payDay)
	salaryDate = checkWeekday(salaryDate)
	daysUntilSalary := upToDays(currentDate, salaryDate)
	return daysUntilSalary, salaryDate.Format(time.RFC850[:18])
}

func whenSalaryYear(payDay int) []string {
	_, salaryDate := getAPIDates(payDay)
	salaryDate = checkWeekday(salaryDate)
	salaryDates := getYearlySalaryDates(salaryDate)
	return salaryDates
}

//func handleQuery(w http.ResponseWriter, rq string) {
//	q, _ := url.ParseQuery(rq)
//	if _, ok := q["pay-day"]; len(q) > 1 || !ok {
//		w.Write([]byte("Invalid query"))
//		return
//	}
//}

// check if there are multiple queries as well - should it be invalidated?
// check if it's bigger then the last day of the month, not 31.
func (app *application) getPayDate(w http.ResponseWriter, r *http.Request) (int, bool) {
	pd, err := strconv.Atoi(r.URL.Query().Get("pay-day"))

	if err != nil || pd < 1 || pd > 31 {
		app.clientError(w, 400)
		return 0, false
	}
	return pd, true
}

func Date(year int, month time.Month, day int) time.Time {
	return time.Date(year, month, day, 0, 0, 0, 0, time.UTC)
}

func getAPIDates(payDay int) (currentDate time.Time, salaryDate time.Time) {
	y, m, d := time.Now().Date()
	currentDate = Date(y, m, d)
	salaryDate = Date(y, m, payDay)
	if payDay < currentDate.Day() {
		salaryDate = Date(y, m+1, payDay)
	}
	return currentDate, salaryDate
}

func getYearlySalaryDates(salaryDate time.Time) []string {
	var salaryDates []string
	y, m, d := salaryDate.Date()

	for m = salaryDate.Month(); m <= 12; m++ {
		salaryDate = Date(y, m, d)
		salaryDate = checkWeekday(salaryDate)
		salaryDates = append(salaryDates, salaryDate.Format(time.RFC850[:17]))
	}

	return salaryDates
}

// upToDays calculates the nr. of days until the salaryDate.
func upToDays(currentDate time.Time, salaryDate time.Time) int {
	days := salaryDate.Sub(currentDate).Hours() / 24
	return int(days)
}

// checkWeekday verifies if the salaryDate is Sunday/Saturday - if that's the case,
// it returns the next friday date.
func checkWeekday(salaryDate time.Time) time.Time {
	y, m, d := salaryDate.Date()

	switch salaryDate.Weekday() {
	case 0:
		salaryDate = Date(y, m, d+5)
	case 6:
		salaryDate = Date(y, m, d+6)
	}
	return salaryDate
}
