package signin

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/cgxarrie-go/signin/pkg/util"
)

type AttendanceRequest struct {
	From time.Time
	To   time.Time
}

type AttendanceResponseData struct {
	Visits   []Visit   `json:"visits,omitempty"`
	Bookings []Booking `json:"bookings,omitempty"`
}

type Visit struct {
	ID          int64  `json:"id"`
	SiteID      int64  `json:"siteId"`
	SiteName    string `json:"siteName"`
	Timezone    string `json:"timezone"`
	InDateTime  string `json:"inDateTime"`
	OutDateTime string `json:"outDateTime"`
}

type Booking struct {
	ID        int64   `json:"id"`
	SiteID    int64   `json:"siteId"`
	SiteName  string  `json:"siteName"`
	Timezone  string  `json:"timezone"`
	Space     Space   `json:"space"`
	Occupancy int     `json:"occupancy"`
	StartDate string  `json:"startDate"`
	EndDate   string  `json:"endDate"`
	Note      *string `json:"note"`
}

type Space struct {
	ID             string    `json:"id"`
	Name           string    `json:"name"`
	Description    *string   `json:"description"`
	PhotoFull      *string   `json:"photoFull"`
	PhotoCover     *string   `json:"photoCover"`
	PhotoThumbnail *string   `json:"photoThumbnail"`
	Category       string    `json:"category"`
	Tags           *[]string `json:"tags"`
	Capacity       int       `json:"capacity"`
	Disabled       bool      `json:"disabled"`
	Zones          []Zone    `json:"zones"`
}

type Zone struct {
	Capacity    int     `json:"capacity"`
	Description *string `json:"description"`
	ID          string  `json:"id"`
	Name        string  `json:"name"`
	Type        string  `json:"type"`
}

// Top-level type for unmarshalling the whole response:
type AttendanceResponse map[string]AttendanceResponseData

// ListFreeSpaces.
func (c client) Attendance(ctx context.Context, req AttendanceRequest) (
	resp AttendanceResponse, err error) {

	format := "%d-%d-%dT%d:%d:%d"
	from := fmt.Sprintf(format, req.From.Year(), req.From.Month(), req.From.Day(), req.From.Hour(), req.From.Minute(), req.From.Second())
	to := fmt.Sprintf(format, req.To.Year(), req.To.Month(), req.To.Day(), req.To.Hour(), req.To.Minute(), req.To.Second())

	url := fmt.Sprintf("%s?startDate=%s&endDate=%s", calendarUrl, from, to)

	err = util.Request(req).
		WithURL(url).
		WithHTTPMethod(http.MethodGet).
		WithContext(ctx).
		WithResponseBody(&resp).
		WithSigninHeders(c.bearer).
		WithExpectedStatusCode(http.StatusOK).
		Do()

	if err != nil {
		return resp, fmt.Errorf("calling calendar: %w", err)
	}

	return resp, nil
}
