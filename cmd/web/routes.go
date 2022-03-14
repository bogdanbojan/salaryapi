package main

import "net/http"

func (app *application) routes() http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("/salary/how-much", app.howMuch)
	mux.HandleFunc("/salary/list-how-many", app.howMany)
	return mux
}
