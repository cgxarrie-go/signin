package service

import (
	"context"
	"fmt"
	"time"

	"github.com/cgxarrie-go/signin/pkg/signin"
	"github.com/cgxarrie-go/signin/pkg/util"
)

// ListBookingsquest.
type ListBookingsquest struct {
	EndDate string
}

// ListBookingsResponse .
type ListBookingsResponse []ListBookingsResponseItem

type ListBookingsResponseItem struct {
	ID       int
	Date     time.Time
	DeskName string
	SiteName string
}

// ListBookings list my bookings from today on
func (s service) ListBookings(ctx context.Context, req ListBookingsquest) (
	resp ListBookingsResponse, err error) {

	now := time.Now()

	limit, err := util.DateFromString(req.EndDate)
	if err != nil {
		return resp, fmt.Errorf("invalid date: %w", err)
	}

	startDate := time.Date(now.Year(), now.Month(), now.Day(), 22, 0, 0, 0,
		now.Location())
	isLastLoop := false
	for {

		endDate := startDate.AddDate(0, 1, 0)
		if endDate.After(limit) {
			endDate = limit
			isLastLoop = true
		}

		items, err := s.getListBookingItems(ctx, startDate, endDate)
		if err != nil {
			return resp, fmt.Errorf("gettig items from %s to %s",
				startDate, endDate)

		}
		resp = append(resp, items...)
		if isLastLoop {
			break
		}

		startDate = endDate.AddDate(0, 0, 1)
	}

	return resp, nil
}

func (s service) getListBookingItems(ctx context.Context, from, to time.Time) (
	resp []ListBookingsResponseItem, err error) {

	clientRequest := signin.ListBookingsRequest{
		From: from,
		To:   to,
	}
	clientResponse, err := s.signinClient.ListBookings(ctx, clientRequest)
	if err != nil {
		return resp, fmt.Errorf("calling signin client: %w", err)
	}

	for _, v := range clientResponse {
		r := ListBookingsResponseItem{
			ID:       v.ID,
			Date:     v.Date,
			DeskName: v.Space.Name,
			SiteName: v.Space.Zones[0].Name,
		}

		resp = append(resp, r)
	}

	return
}
