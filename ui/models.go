package ui

import "time"

var (
	idColLength   int = 10
	dateColLength int = 25
	deskColLength int = 25
	siteColLength int = 10
	weekColLength int = 10

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

type AttendanceWeek struct {
	RelativeWeek        int
	WeekStartDate       time.Time
	WeekEndDate         time.Time
	Week                string
	WorkingDays         int
	Bookings            int
	Visits              int
	VisitsPerWorkingDay float64
}

type AttendanceSummary struct {
	WorkingDays   int
	Visits        int
	AvgOfficeTime float64
}

type Attendance struct {
	Weeks   map[int]AttendanceWeek
	Summary AttendanceSummary
}

type AttendanceFreeDay struct {
	Date   time.Time
	Reason string
}
type AttendanceFreeDays struct {
	FreeDays []AttendanceFreeDay
}
