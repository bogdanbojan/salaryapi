package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

var app = &application{ // synchronization prob?
	errorLog: nil,
	infoLog:  nil,
}

func TestHowMuch(t *testing.T) {
	t.Parallel() // ?
	for i, urc := range URLRequestsCases.howMuchRequests {
		rr := httptest.NewRecorder() // ?
		r, err := http.NewRequest(URLRequestsCases.methodRequests[i], urc, nil)
		if err != nil {
			t.Fatal(err)
		}
		app.howMuch(rr, r)
		rs := rr.Result()
		assertStatusCode(t, rs.StatusCode, URLResponsesCases.howMuchResponses[i])
		err = rs.Body.Close()
		if err != nil {
			return
		}
	}
}

func TestHowMany(t *testing.T) {
	t.Parallel()
	for i, urc := range URLRequestsCases.howManyRequests {
		rr := httptest.NewRecorder() // ?
		r, err := http.NewRequest(URLRequestsCases.methodRequests[i], urc, nil)
		if err != nil {
			t.Fatal(err)
		}
		app.howMany(rr, r)
		rs := rr.Result()
		assertStatusCode(t, rs.StatusCode, URLResponsesCases.howManyResponses[i])
		err = rs.Body.Close()
		if err != nil {
			return
		}
	}
}

func assertStatusCode(t testing.TB, got, want int) {
	t.Helper()
	if got != want {
		t.Errorf("got %d, want %d", got, want)
	}

}

var URLRequestsCases = struct {
	methodRequests  []string
	howMuchRequests []string
	howManyRequests []string
}{
	methodRequests: []string{
		"GET",
		"GET",
		"UPDATE",
		"GET",
		"POST",
	},
	howMuchRequests: []string{
		"/salary/how-much?pay-day=20",
		"/salary/how-much?pay-day=-20",
		"/salary/how-much?pay-day=40",
		"/salary/how-much?pay-day=0",
		"/annual/festival/of/not/working",
	},
	howManyRequests: []string{
		"/salary/list-how-many?pay-day=20",
		"/salary/list-how-many?pay-day=-20",
		"/salary/list-how-many?pay-day=40",
		"/salary/list-how-many?pay-day=0",
		"/annual/festival/of/not/working",
	},
}
var URLResponsesCases = struct {
	howMuchResponses []int
	howManyResponses []int
}{
	howMuchResponses: []int{
		http.StatusOK,
		http.StatusBadRequest,
		http.StatusMethodNotAllowed,
		http.StatusBadRequest,
		http.StatusMethodNotAllowed,
	},
	howManyResponses: []int{
		http.StatusOK,
		http.StatusBadRequest,
		http.StatusMethodNotAllowed,
		http.StatusBadRequest,
		http.StatusMethodNotAllowed,
	},
}

func TestGetPayDate(t *testing.T) {
	for i, urc := range URLRequestsCases.howManyRequests {
		rr := httptest.NewRecorder()
		r, err := http.NewRequest("GET", urc, nil)
		if err != nil {
			t.Fatal(err)
		}
		payDay, flag := app.getPayDay(rr, r)
		assertPayDayResponse(t, payDay, flag, payDayResponsesCases.payDay[i], payDayResponsesCases.flag[i])
	}
}

// tuple? Pair struct{a,b interface{}}
func assertPayDayResponse(t testing.TB, gotInt int, gotBool bool, wantInt int, wantBool bool) {
	t.Helper()
	if gotInt != wantInt {
		if gotInt != wantInt {
			t.Errorf("got %d, want %d", gotInt, wantInt)
		}
	}
	if gotBool != wantBool {
		if gotBool != wantBool {
			t.Errorf("got %t, want %t", gotBool, wantBool)
		}
	}
}

var payDayResponsesCases = struct {
	payDay []int
	flag   []bool
}{
	payDay: []int{
		20,
		0,
		0,
		0,
		0,
	},
	flag: []bool{
		true,
		false,
		false,
		false,
		false,
	},
}
