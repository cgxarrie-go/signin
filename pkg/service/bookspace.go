package service

import (
	"context"
	"fmt"
	"strconv"
	"time"

	"github.com/cgxarrie-go/signin/config"
	"github.com/cgxarrie-go/signin/pkg/signin"
	"github.com/cgxarrie-go/signin/pkg/util"
)

// BookSpaceRequest .
type BookSpaceRequest struct {
	DeskNumber string
	Date       string
}

// BookSpaceResponse .
type BookSpaceResponse struct {
	ID       int
	DeskID   string
	DeskName string
	ZoneName string
	Date     string
}

// BookSpace book a desk
func (s service) BookSpace(ctx context.Context, req BookSpaceRequest) (resp BookSpaceResponse,
	err error) {

	deskNum, err := strconv.Atoi(req.DeskNumber)
	if err != nil {
		return resp, fmt.Errorf("invalid desk number : %s", req.DeskNumber)
	}

	date, err := util.DateFromString(req.Date)
	if err != nil {
		return resp, err
	}

	endDate := date.Add(21 * time.Hour).
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
		StartDate: startDate,
		EndDate:   endDate,
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
		Date:     clientResponse.EndDate.Format("2006-01-02"),
	}

	return resp, nil
}
