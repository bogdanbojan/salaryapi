package main

import (
	"net/http"
	"net/url"

	"strconv"
)

func (app *application) howMuch(w http.ResponseWriter, r *http.Request) {
	if handleOtherThanGET(w, r) {
		return
	}
	handleQuery(w, r.URL.RawQuery)

	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte("hm"))
}

func (app *application) howMany(w http.ResponseWriter, r *http.Request) {
	if handleOtherThanGET(w, r) {
		return
	}
	handleQuery(w, r.URL.RawQuery)

	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte("Creating something.."))
}

func handleOtherThanGET(w http.ResponseWriter, r *http.Request) bool {
	if r.Method != "GET" {
		w.Header().Set("Allow", "GET")
		http.Error(w, "Method Not Allowed", 405)
		return true
	}
	return false
}

func handleQuery(w http.ResponseWriter, rq string) {
	q, _ := url.ParseQuery(rq)
	if _, ok := q["pay-day"]; len(q) > 1 || !ok {
		w.Write([]byte("Invalid query"))
		return
	}
}

func verifyPayDate(w http.ResponseWriter, r *http.Request) {
	pd, err := strconv.Atoi(r.URL.Query().Get("pay-date"))
	if err != nil || pd < 1 {
		http.NotFound(w, r)
		return
	}
}

func verifyPeriod(w http.ResponseWriter, r *http.Request) {

}
