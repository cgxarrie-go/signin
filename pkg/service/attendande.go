package service

import (
	"context"
	"fmt"
	"time"

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
	RelativeWeek     int
	WeekStartDate    time.Time
	WeekEndDate      time.Time
	Week             string
	NumberOfBookings int
	NumberOfVisits   int
}

type AttendanceSummary struct {
	AverageVisitsPerWeek   float64
	AverageBookingsPerWeek float64
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
			AverageVisitsPerWeek:   0,
			AverageBookingsPerWeek: 0,
		},
	}

	ttlVisits := 0
	ttlBookings := 0
	ttlitems := 0

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
		for d := 0; d < 7; d++ {
			day := weekStart.AddDate(0, 0, d)
			dayKey := day.Format("2006-01-02")
			if data, ok := clientResponse[dayKey]; ok {
				bookings += len(data.Bookings)
				visits += len(data.Visits)
			}
		}
		item.NumberOfBookings = bookings
		ttlBookings += bookings
		item.NumberOfVisits = visits
		ttlVisits += visits
		resp.Items[-i] = item
		ttlitems++
	}

	if ttlitems > 0 {
		resp.Summary.AverageVisitsPerWeek = float64(ttlVisits) / float64(ttlitems)
		resp.Summary.AverageBookingsPerWeek = float64(ttlBookings) / float64(ttlitems)
	}

	return resp, nil
}
