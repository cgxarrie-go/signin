package service

import (
	"context"
	"fmt"
	"strconv"
	"time"

	"github.com/cgxarrie-go/signin/config"
	"github.com/cgxarrie-go/signin/pkg/signin"
)

// BookSpaceRequest .
type BookSpaceRequest struct {
	DeskNumber string
	Items      int
	Dates      []time.Time
}

// BookSpaceResponse .
type BookSpaceResponse struct {
	ID       int
	DeskID   string
	DeskName string
	ZoneName string
	Date     time.Time
}

// BookSpace book a desk
func (s service) BookSpace(ctx context.Context, req BookSpaceRequest) (
	resp []BookSpaceResponse, err error) {

	deskNum, err := strconv.Atoi(req.DeskNumber)
	if err != nil {
		return resp, fmt.Errorf("invalid desk number : %s", req.DeskNumber)
	}

	for i := 0; i < req.Items-1; i++ {
		d := req.Dates[0].AddDate(0, 0, i+1)
		req.Dates = append(req.Dates, d)
	}

	for _, reqDate := range req.Dates {
		if err != nil {
			return resp, err
		}

		r, err := s.bookSpaceForOneDate(ctx, deskNum, reqDate)
		if err != nil {
			r.Date = reqDate
			r.DeskName = req.DeskNumber
			r.ZoneName = err.Error()
		}
		resp = append(resp, r)
	}

	return resp, nil
}

func (s service) bookSpaceForOneDate(ctx context.Context, deskNum int,
	date time.Time) (resp BookSpaceResponse, err error) {

	endDate := date.
		Add(22 * time.Hour).
		Add(59 * time.Minute).
		Add(59 * time.Second)

	startDate := endDate.Add(-24 * time.Hour).
		Add(1 * time.Second)

	spaceID, ok := config.Instance().Desks[deskNum]
	if !ok {
		err = s.GetSpaceIDs(ctx)
		if err != nil {
			return resp, fmt.Errorf("getting desk IDs: %w", err)
		}
		spaceID, ok = config.Instance().Desks[deskNum]
		if !ok {
			return resp, fmt.Errorf("desk number %d not in Desk Map", deskNum)
		}
	}

	clientRequest := signin.BookSpaceRequest{
		SiteID:    34098,
		SpaceID:   spaceID,
		Occupancy: 1,
		StartDate: startDate.Format("2006-01-02T15:04:05.000Z"),
		EndDate:   endDate.Format("2006-01-02T15:04:05.000Z"),
		Note:      "",
		SendMail:  true,
	}

	clientResponse, err := s.signinClient.BookSpace(ctx, clientRequest)
	if err != nil {
		return resp, fmt.Errorf("calling signin client: %w", err)
	}

	resp = BookSpaceResponse{
		ID:       clientResponse.ID,
		DeskID:   clientResponse.Space.ID,
		DeskName: clientResponse.Space.Name,
		ZoneName: clientResponse.Space.Zones[0].Name,
		Date:     clientResponse.EndDate,
	}

	return resp, nil
}
