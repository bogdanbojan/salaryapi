package main

import (
	"net/http"
	"strconv"
)

func (app *application) howMuch(w http.ResponseWriter, r *http.Request) {
	handleOtherRequestMethods(w, r)
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte("Creating something.."))

}

func (app *application) howMany(w http.ResponseWriter, r *http.Request) {
	handleOtherRequestMethods(w, r)
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte("Creating something.."))

}

func handleOtherRequestMethods(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		w.Header().Set("Allow", "GET")
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, "Method Not Allowed", 405)
		return
	}
	return
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
