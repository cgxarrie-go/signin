package service

import (
	"context"
	"fmt"
	"time"

	"github.com/cgxarrie-go/signin/config"
	"github.com/cgxarrie-go/signin/pkg/signin"
)

// AttendanceRequest .
type AttendanceRequest struct {
	NumberOfWeeks int
}

// AttendanceResponse .
type AttendanceResponse struct {
	Items   map[int]AttendanceResponseItem
	Summary AttendanceSummary
}

type AttendanceResponseItem struct {
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
	Bookings      int
	AvgOfficeTime float64
}

// BookSpace book a desk
func (s service) Attendance(ctx context.Context, req AttendanceRequest) (
	resp AttendanceResponse, err error) {

	now := time.Now()
	weekday := int(now.Weekday())
	// In Go, Sunday == 0, so adjust to Monday == 0
	if weekday == 0 {
		weekday = 7
	}
	currentMonday := now.AddDate(0, 0, -weekday+1).Truncate(24 * time.Hour)
	currentDay := now.Truncate(24 * time.Hour)

	startDate := currentMonday.AddDate(0, 0, -7*req.NumberOfWeeks)
	endDate := currentDay.AddDate(0, 0, 1).Add(-1 * time.Second)

	clientRequest := signin.AttendanceRequest{
		From: startDate,
		To:   endDate,
	}

	clientResponse, err := s.signinClient.Attendance(ctx, clientRequest)
	if err != nil {
		return resp, fmt.Errorf("calling signin client: %w", err)
	}

	resp = AttendanceResponse{
		Items: make(map[int]AttendanceResponseItem),
		Summary: AttendanceSummary{
			WorkingDays:   0,
			Visits:        0,
			Bookings:      0,
			AvgOfficeTime: 0,
		},
	}

	ttlWorkingDays := 0
	ttlVisits := 0
	ttlBookings := 0

	for i := 0; i < req.NumberOfWeeks; i++ {
		weekStart := currentMonday.AddDate(0, 0, -7*(req.NumberOfWeeks-i-1))
		weekEnd := weekStart.AddDate(0, 0, 6)
		wy, wn := weekStart.ISOWeek()

		item := AttendanceResponseItem{
			RelativeWeek:  i - req.NumberOfWeeks + 1,
			WeekStartDate: weekStart,
			WeekEndDate:   weekEnd,
			Week:          fmt.Sprintf("%04d-%02d", wy, wn),
		}

		bookings := 0
		visits := 0
		workingDays := 0
		for d := 0; d < 7; d++ {
			day := weekStart.AddDate(0, 0, d)
			if s.isWorkingDay(day) {
				workingDays++
			}

			dayKey := day.Format("2006-01-02")
			if data, ok := clientResponse[dayKey]; ok {
				bookings += len(data.Bookings)
				visits += len(data.Visits)
			}
		}
		item.Bookings = bookings
		item.Visits = visits
		item.WorkingDays = workingDays
		if workingDays > 0 {
			item.VisitsPerWorkingDay = float64(visits) / float64(workingDays)
		}
		resp.Items[-i] = item
		ttlVisits += visits
		ttlBookings += bookings
		ttlWorkingDays += workingDays
	}

	if len(resp.Items) > 0 {
		resp.Summary.WorkingDays = ttlWorkingDays
		resp.Summary.Visits = ttlVisits
		resp.Summary.Bookings = ttlBookings
		resp.Summary.AvgOfficeTime = float64(ttlVisits) / float64(ttlWorkingDays)
	}

	return resp, nil
}

func (s service) isWorkingDay(day time.Time) bool {
	// Check if the day is a weekend
	if day.Weekday() == time.Saturday || day.Weekday() == time.Sunday {
		return false
	}

	_, ok := config.Instance().AttendanceFreeDays[day.Format("2006-01-02")]
	if ok {
		return false
	}

	return true
}
