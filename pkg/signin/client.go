package signin

import (
	"context"
	"fmt"
)

type Client interface {
	BookSpace(ctx context.Context, req BookSpaceRequest) (
		resp BookSpaceRespnse, err error)
	CancelBooking(ctx context.Context, req CancelBookingRequest) (
		resp CancelBookingRespnse, err error)
	ListBookings(ctx context.Context, req ListBookingsRequest) (
		resp ListBookingsRespnse, err error)
	ListFreeSpaces(ctx context.Context, req ListFreeSpacesRequest) (
		resp ListFreeSpacesResponse, err error)
}

type client struct {
	bearer string
}

const (
	baseURL string = "https://backend.signinapp.com"
)

var (
	bookSpaceUrl      string = fmt.Sprintf("%s/api/mobile/spaces/bookings", baseURL)
	cancelBookingUrl  string = fmt.Sprintf("%s/api/mobile/spaces/bookings", baseURL)
	listBookingsUrl   string = fmt.Sprintf("%s/api/mobile/spaces/bookings", baseURL)
	listFreeSpacesUrl string = fmt.Sprintf("%s/api/mobile/spaces/34098/"+
		"search?occupancy=1", baseURL)
)

func NewClient(bearer string) Client {
	return client{
		bearer: bearer,
	}
}
