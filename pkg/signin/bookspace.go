package signin

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/cgxarrie-go/signin/pkg/util"
)

type BookSpaceRequest struct {
	SiteID    int    `json:"site_id"`
	SpaceID   string `json:"space_id"`
	Occupancy int    `json:"occupancy"`
	StartDate string `json:"start_date"`
	EndDate   string `json:"end_date"`
	Note      string `json:"note"`
	SendMail  bool   `json:"send_confirmation_email"`
}

type BookSpaceRespnse struct {
	Item
	Occupancy int       `json:"occupancy"`
	StartDate time.Time `json:"start_date"`
	EndDate   time.Time `json:"end_date"`
	Note      string    `json:"note"`
}

// BookSpace .
func (c client) BookSpace(ctx context.Context, req BookSpaceRequest) (
	resp BookSpaceRespnse, err error) {

	err = util.Request(req).
		WithURL(bookSpaceUrl).
		WithHTTPMethod(http.MethodPost).
		WithContext(ctx).
		WithSigninHeders(c.bearer).
		WithResponseBody(&resp).
		WithExpectedStatusCode(http.StatusCreated).
		Do()

	if err != nil {
		return resp, fmt.Errorf("calling book space: %w", err)
	}

	return resp, nil
}
