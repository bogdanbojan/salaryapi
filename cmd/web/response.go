package main

import (
	"encoding/json"
	"net/http"
)

// Monthly and Yearly are the JSON response structs.
type Monthly struct {
	NrOfDays int    `json:"nr-of-days"`
	Date     string `json:"date"`
}

type Yearly struct {
	Dates []string `json:"dates"`
}

// writeJSON writes a JSON response to http requests. It can work with multiple data types and can also write a header.
// The latter one is optional. The for loop is safe since we can range over nil map. Go does not throw an error if you
// try to range over (or generally, read from) a nil map.
func (app *application) writeJSON(w http.ResponseWriter, status int, data interface{}, headers http.Header) error {
	js, err := json.MarshalIndent(data, "", "\t")
	if err != nil {
		return err
	}

	js = append(js, '\n')

	for key, value := range headers {
		w.Header()[key] = value
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_, err = w.Write(js)
	if err != nil {
		return err
	}
	return nil
}
