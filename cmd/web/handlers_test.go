package main

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHowMuch(t *testing.T) {
	t.Parallel()

	rr := httptest.NewRecorder()

	r, err := http.NewRequest("GET", "/salary/how-much?pay-day=20", nil)
	if err != nil {
		t.Fatal(err)
	}

	app := application{
		errorLog: nil,
		infoLog:  nil,
	}
	app.howMuch(rr, r)
	rs := rr.Result()
	if rs.StatusCode != http.StatusOK {
		t.Errorf("want %d got %d", http.StatusOK, rs.StatusCode)
	}
	defer rs.Body.Close()
	_, err = ioutil.ReadAll(rs.Body)
	if err != nil {
		t.Fatal(err)
	}
}
