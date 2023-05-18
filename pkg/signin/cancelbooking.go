package signin

import (
	"context"
	"fmt"
	"net/http"

	"github.com/cgxarrie-go/signin/pkg/util"
)

type CancelBookingRequest struct {
	BookingID int
}
type CancelBookingRespnse struct{}

// CancelBooking .
func (c client) CancelBooking(ctx context.Context, req CancelBookingRequest) (
	resp CancelBookingRespnse, err error) {

	url := fmt.Sprintf("%s/%d", cancelBookingUrl, req.BookingID)
	err = util.Request(req).
		WithURL(url).
		WithHTTPMethod(http.MethodDelete).
		WithContext(ctx).
		WithSigninHeders(c.bearer).
		WithExpectedStatusCode(http.StatusNoContent).
		Do()

	if err != nil {
		return resp, fmt.Errorf("calling cancel booking: %w", err)
	}

	return resp, nil
}
