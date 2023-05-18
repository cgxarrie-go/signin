package signin

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/cgxarrie-go/signin/pkg/util"
)

type ListBookingsRequest struct {
	From time.Time
	To   time.Time
}
type ListBookingsRespnse []ListBookingsRespnseItem

type ListBookingsRespnseItem struct {
	Item
	Date time.Time `json:"end_date"`
}

// ListBookings .
func (c client) ListBookings(ctx context.Context, req ListBookingsRequest) (
	resp ListBookingsRespnse, err error) {

	url := fmt.Sprintf("%s?start_date=%s&end_date=%s", listBookingsUrl,
		req.From.Format("2006-01-02T04:05:06Z"),
		req.To.Format("2006-01-02T04:05:06Z"))

	err = util.Request(req).
		WithURL(url).
		WithHTTPMethod(http.MethodGet).
		WithContext(ctx).
		WithResponseBody(&resp).
		WithSigninHeders(c.bearer).
		WithExpectedStatusCode(http.StatusOK).
		Do()

	if err != nil {
		return resp, fmt.Errorf("calling cancel booking: %w", err)
	}

	return resp, nil
}
