package main

import (
	"net/http"
	"net/url"
	"time"

	"strconv"
)

func (app *application) howMuch(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		pd, ok := getPayDate(w, r)
		if !ok {
			return
		}
		w.Header().Set("Content-Type", "application/json")
		days, date := whenSalaryMonth(pd)
		w.Write([]byte("You have " + strconv.Itoa(days) + " more days till your payday."))
		w.Write([]byte("Paydate: " + date))

	default:
		w.Header().Set("Allow", "GET")
		http.Error(w, "Method Not Allowed", 405)
	}

}

func (app *application) howMany(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		pd, ok := getPayDate(w, r)
		if !ok {
			return
		}
		w.Header().Set("Content-Type", "application/json")
		dates := whenSalaryYear(pd)
		for _, date := range dates {
			w.Write([]byte(date))
			w.Write([]byte("\n"))
		}
	default:
		w.Header().Set("Allow", "GET")
		http.Error(w, "Method Not Allowed", 405)
	}
}

func Date(year int, month time.Month, day int) time.Time {
	return time.Date(year, month, day, 0, 0, 0, 0, time.UTC)

}

func whenSalaryMonth(day int) (int, string) {
	y, m, d := time.Now().Date()
	currentDate := Date(y, m, d)
	salaryDate := Date(y, m, day)
	if day < currentDate.Day() {
		salaryDate = Date(y, m+1, day)
	}

	switch salaryDate.Day() {
	case 0:
		salaryDate = Date(y, m, day+5)
	case 6:
		salaryDate = Date(y, m, day+6)
	}

	days := salaryDate.Sub(currentDate).Hours() / 24
	date := salaryDate.String()
	return int(days), date[:10]
}

func whenSalaryYearHelper(salaryDate time.Time) string {
	y, m, d := salaryDate.Date()
	switch salaryDate.Day() {
	case 0:
		salaryDate = Date(y, m, d+5)
	case 6:
		salaryDate = Date(y, m, d+6)
	}

	return salaryDate.String()[:10]

}

// TODO: year return not working properly
func whenSalaryYear(day int) []string {
	y, m, d := time.Now().Date()
	currentDate := Date(y, m, d)
	salaryDate := Date(y, m, day)

	if salaryDate.Day() < currentDate.Day() {
		salaryDate = Date(y, m+1, day)
	}

	var salariesDates []string

	for m = salaryDate.Month(); m <= 12; m++ {
		salaryDate = Date(y, m, day)
		date := whenSalaryYearHelper(salaryDate)
		salariesDates = append(salariesDates, date)

		//currentDate = Date(y, m+1, d)
	}

	return salariesDates
}

func handleQuery(w http.ResponseWriter, rq string) {
	q, _ := url.ParseQuery(rq)
	if _, ok := q["pay-day"]; len(q) > 1 || !ok {
		w.Write([]byte("Invalid query"))
		return
	}
}

// check if there are multiple queries as well - should it be invalidated?
func getPayDate(w http.ResponseWriter, r *http.Request) (int, bool) {
	pd, err := strconv.Atoi(r.URL.Query().Get("pay-day"))
	if err != nil || pd < 1 || pd > 31 {
		http.NotFound(w, r)
		return 0, false
	}
	return pd, true
}
