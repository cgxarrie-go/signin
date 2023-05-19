package ui

import "time"

var (
	idColLength   int = 10
	dateColLength int = 25
	deskColLength int = 25
	siteColLength int = 10

	dateFormat = "2006-01-02"
)

type Booking struct {
	ID int
	Desk
	Date time.Time
}

type Desk struct {
	ID       string
	Name     string
	ZoneName string
}
