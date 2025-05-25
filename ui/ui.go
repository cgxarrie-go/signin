package ui

import (
	"fmt"
	"strings"
	"sync"
)

var uiLock = &sync.Mutex{}

var uiSingleton *ui

// Instance
func Instance() *ui {
	if uiSingleton == nil {
		uiLock.Lock()
		defer uiLock.Unlock()
		uiSingleton = &ui{}
	}

	return uiSingleton
}

// ui
type ui struct {
}

func (ui ui) PrintBooking(booking Booking) {
	bookings := []Booking{booking}
	ui.PrintBookings(bookings)
}

func (ui ui) PrintBookings(bookings []Booking) {

	titleFormat := "%" + fmt.Sprintf("%d", idColLength) + "s " +
		"| %-" + fmt.Sprintf("%d", dateColLength) + "s " +
		"| %-" + fmt.Sprintf("%d", deskColLength) + "s " +
		"| %-" + fmt.Sprintf("%d", siteColLength) + "s "

	title := fmt.Sprintf(titleFormat, "ID", "Date", "Zone", "Desk")
	line := strings.Repeat("-", len(title)+5)

	rowFormat := "%" + fmt.Sprintf("%d", idColLength) + "d " +
		"| %-" + fmt.Sprintf("%d", dateColLength) + "s " +
		"| %-" + fmt.Sprintf("%d", deskColLength) + "s " +
		"| %-" + fmt.Sprintf("%d", siteColLength) + "s "

	fmt.Println(title)
	fmt.Println(line)

	for _, b := range bookings {
		info := fmt.Sprintf(rowFormat, b.ID, b.Date.Format(dateFormat),
			b.ZoneName, b.Name)

		fmt.Println(info)
	}

	fmt.Println(line)

}

func (ui ui) PrintDesks(desks []Desk) {

	titleFormat := "%" + fmt.Sprintf("%d", idColLength) + "s " +
		"| %" + fmt.Sprintf("%d", deskColLength) + "s " +
		"| %" + fmt.Sprintf("%d", siteColLength) + "s "

	title := fmt.Sprintf(titleFormat, "ID", "Zone", "Desk")
	line := strings.Repeat("-", len(title)+5)

	rowFormat := "%" + fmt.Sprintf("%d", idColLength) + "s " +
		"| %" + fmt.Sprintf("%d", deskColLength) + "s " +
		"| %" + fmt.Sprintf("%d", siteColLength) + "s "
	fmt.Println(title)
	fmt.Println(line)

	for _, b := range desks {
		info := fmt.Sprintf(rowFormat, b.ID, b.ZoneName, b.Name)
		fmt.Println(info)
	}

	fmt.Println(line)
}

func (ui ui) PrintAttendance(attendance Attendance) {

	titleFormat := "%" + fmt.Sprintf("%d", idColLength) + "s " +
		"| %-" + fmt.Sprintf("%d", weekColLength) + "s " +
		"| %-" + fmt.Sprintf("%d", dateColLength) + "s " +
		"| %-" + fmt.Sprintf("%d", dateColLength) + "s " +
		"| %-" + fmt.Sprintf("%d", idColLength) + "s " +
		"| %-" + fmt.Sprintf("%d", idColLength) + "s " +
		"| %-" + fmt.Sprintf("%d", idColLength) + "s " +
		"| %-" + fmt.Sprintf("%d", idColLength) + "s "

	title := fmt.Sprintf(titleFormat, "", "Week", "Start", "End",
		"WrkDays", "Bookings", "Visits", "Visits/WkDay")
	line := strings.Repeat("-", len(title)+5)

	rowFormat := "%" + fmt.Sprintf("%d", idColLength) + "d " +
		"| %-" + fmt.Sprintf("%d", weekColLength) + "s " +
		"| %-" + fmt.Sprintf("%d", dateColLength) + "v " +
		"| %-" + fmt.Sprintf("%d", dateColLength) + "v " +
		"| %" + fmt.Sprintf("%d", idColLength) + "d " +
		"| %" + fmt.Sprintf("%d", idColLength) + "d " +
		"| %" + fmt.Sprintf("%d", idColLength) + "d " +
		"| %" + fmt.Sprintf("%d", idColLength) + ".2f%%"

	fmt.Println(title)
	fmt.Println(line)

	for i := 0; i < len(attendance.Weeks); i++ {
		b := attendance.Weeks[i]
		info := fmt.Sprintf(rowFormat, b.RelativeWeek, b.Week,
			b.WeekStartDate.Format(dateFormat),
			b.WeekEndDate.Format(dateFormat),
			b.WorkingDays, b.Bookings, b.Visits, b.VisitsPerWorkingDay)
		fmt.Println(info)
	}

	fmt.Println(line)
	fmt.Printf("Total Working Days: %d\n", attendance.Summary.WorkingDays)
	fmt.Printf("Total Visits: %d\n", attendance.Summary.Visits)
	fmt.Printf("Avg Office Time: %.2f%%\n", attendance.Summary.AvgOfficeTime)
	fmt.Println(line)
}

func (ui ui) PrintAttendanceFreeDays(days AttendanceFreeDays) {
	titleFormat := "%" + fmt.Sprintf("%d", dateColLength) + "s " +
		"| %-" + fmt.Sprintf("%d", idColLength) + "s "

	title := fmt.Sprintf(titleFormat, "Date", "Reason")
	line := strings.Repeat("-", len(title)+5)

	rowFormat := "%" + fmt.Sprintf("%d", dateColLength) + "v " +
		"| %-" + fmt.Sprintf("%d", idColLength) + "s "

	fmt.Println(title)
	fmt.Println(line)

	for _, b := range days.FreeDays {
		info := fmt.Sprintf(rowFormat, b.Date.Format(dateFormat), b.Reason)
		fmt.Println(info)
	}

	fmt.Println(line)
}
