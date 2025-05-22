package signin

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/cgxarrie-go/signin/pkg/util"
)

type ListFreeSpacesRequest struct {
	From time.Time
	To   time.Time
}
type ListFreeSpacesResponse []ItemSpace

// ListFreeSpaces.
func (c client) ListFreeSpaces(ctx context.Context, req ListFreeSpacesRequest) (
	resp ListFreeSpacesResponse, err error) {

	format := "%d-%d-%dT%d:%d:%d"
	from := fmt.Sprintf(format, req.From.Year(), req.From.Month(), req.From.Day(), req.From.Hour(), req.From.Minute(), req.From.Second())
	to := fmt.Sprintf(format, req.To.Year(), req.To.Month(), req.To.Day(), req.To.Hour(), req.To.Minute(), req.To.Second())
	url := fmt.Sprintf("%s&start_date=%s&end_date=%s", listFreeSpacesUrl, from, to)

	err = util.Request(req).
		WithURL(url).
		WithHTTPMethod(http.MethodGet).
		WithContext(ctx).
		WithResponseBody(&resp).
		WithSigninHeders(c.bearer).
		WithExpectedStatusCode(http.StatusOK).
		Do()

	if err != nil {
		return resp, fmt.Errorf("calling list free spaces: %w", err)
	}

	return resp, nil
}
