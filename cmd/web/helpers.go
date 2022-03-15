package main

import "time"

const (
	Year = iota
	Month
)

// whenSalaryMonth is a controller which handles the calculation of the next monthly pay date.
func whenSalaryMonth(payDay int) (int, string) {
	currentDate, salaryDate := getAPIDates(payDay)
	salaryDate = checkWeekday(salaryDate)
	daysUntilSalary := upToDays(currentDate, salaryDate)
	return daysUntilSalary, salaryDate.Format(time.RFC850[:18])
}

// whenSalaryYear is a controller which handles the calculation of the next yearly pay dates.
func whenSalaryYear(payDay int) []string {
	_, salaryDate := getAPIDates(payDay)
	salaryDate = checkWeekday(salaryDate)
	salaryDates := getYearlySalaryDates(salaryDate)
	return salaryDates
}

// Date is a factory function that constructs a time.Time value.
func Date(year int, month time.Month, day int) time.Time {
	return time.Date(year, month, day, 0, 0, 0, 0, time.UTC)
}

// getAPIDates returns the next pay date and the current time of your query.
func getAPIDates(payDay int) (currentDate time.Time, salaryDate time.Time) {
	y, m, d := time.Now().Date()
	currentDate = Date(y, m, d)
	salaryDate = Date(y, m, payDay)
	if payDay < currentDate.Day() {
		salaryDate = Date(y, m+1, payDay)
	}
	return currentDate, salaryDate
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

// upToDays calculates the nr. of days until the salaryDate.
func upToDays(currentDate time.Time, salaryDate time.Time) int {
	days := salaryDate.Sub(currentDate).Hours() / 24
	return int(days)
}

// getYearlySalaryDates returns an array of strings with all the dates of
// your pay date in that year formatted to a RFC850 standard (wo time) - day, dd-mmm-yy
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
