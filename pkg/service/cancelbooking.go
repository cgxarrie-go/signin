package service

import (
	"context"
	"fmt"

	"github.com/cgxarrie-go/signin/pkg/signin"
	"github.com/cgxarrie-go/signin/pkg/util"
)

// CancelBookingRequest .
type CancelBookingRequest struct {
	Date string
}

// CancelBooking cancel reservation
func (s service) CancelBooking(ctx context.Context, req CancelBookingRequest) error {

	date, err := util.DateFromString(req.Date)
	if err != nil {
		return fmt.Errorf("invalid date: %w", err)
	}

	items, err := s.getListBookingItems(ctx, date, date)
	if err != nil {
		return fmt.Errorf("getting date bookings: %w", err)
	}

	for _, item := range items {
		clientRequest := signin.CancelBookingRequest{
			BookingID: item.ID,
		}

		_, err = s.signinClient.CancelBooking(ctx, clientRequest)
		if err != nil {
			return fmt.Errorf("calling signin client: %w", err)
		}
	}

	return nil

}
