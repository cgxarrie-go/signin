package util

import (
	"fmt"
	"time"
)

// DateFromString returns a date from a string in format YYYYMMDD
func DateFromString(value string) (date time.Time, err error) {
	date, err = time.Parse("20060102", value)
	if err != nil {
		return date, fmt.Errorf(
			"Invali date : %s. Should be formatted as YYYYMMDD", value)
	}

	return date, nil
}
