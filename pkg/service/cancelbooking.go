package service

import (
	"context"
	"fmt"
	"strconv"

	"github.com/cgxarrie-go/signin/pkg/signin"
)

// CancelBookingRequest .
type CancelBookingRequest struct {
	BookingID string
}

// CancelBooking cancel reservation
func (s service) CancelBooking(ctx context.Context, req CancelBookingRequest) error {

	id, err := strconv.Atoi(req.BookingID)
	if err != nil {
		return fmt.Errorf("invalid booking ID : %s", req.BookingID)
	}

	clientRequest := signin.CancelBookingRequest{
		BookingID: id,
	}

	_, err = s.signinClient.CancelBooking(ctx, clientRequest)
	if err != nil {
		return fmt.Errorf("calling signin client: %w", err)
	}

	return nil

}
