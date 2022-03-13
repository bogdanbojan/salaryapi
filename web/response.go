package main

// Monthly and Yearly are the JSON response structs.
type Monthly struct {
	NrOfDays int
	Date     int
}

type Yearly struct {
	Dates []string
}
