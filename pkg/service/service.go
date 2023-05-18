package service

import (
	"context"

	"github.com/cgxarrie-go/signin/pkg/signin"
)

type Service interface {
	BookSpace(ctx context.Context, req BookSpaceRequest) (
		resp BookSpaceResponse, err error)
	CancelBooking(ctx context.Context, req CancelBookingRequest) error
	ListFreeSpaces(ctx context.Context, req ListFreeSpacesRequest) (
		resp ListFreeSpacesResponse, err error)
	ListBookings(ctx context.Context, req ListBookingsquest) (
		resp ListBookingsResponse, err error)
	GetSpaceIDs(ctx context.Context) (err error)
}

// service
type service struct {
	signinClient signin.Client
}

func New(signinClient signin.Client) Service {
	return service{
		signinClient: signinClient,
	}
}
