package main

import (
	"encoding/json"
	"net/http"
)

// Monthly and Yearly are the JSON response structs.
type Monthly struct {
	NrOfDays int
	Date     string
}

type Yearly struct {
	Dates []string
}

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