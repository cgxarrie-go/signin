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
