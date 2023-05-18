package service

import (
	"context"
	"fmt"
	"time"

	"github.com/cgxarrie-go/signin/pkg/signin"
	"github.com/cgxarrie-go/signin/pkg/util"
)

// ListFreeSpacesRequest.
type ListFreeSpacesRequest struct {
	Date string
}

// ListFreeSpacesResponse .
type ListFreeSpacesResponse []ListFreeSpacesResponseItem

type ListFreeSpacesResponseItem struct {
	ID       string
	DeskName string
	SiteName string
}

// ListFreeSpaces list free desks for a given date
func (s service) ListFreeSpaces(ctx context.Context, req ListFreeSpacesRequest) (
	resp ListFreeSpacesResponse, err error) {

	date, err := util.DateFromString(req.Date)
	if err != nil {
		return resp, err
	}

	endDate := date.Add(21 * time.Hour).Add(59 * time.Minute).Add(59 * time.Second)
	startDate := date.AddDate(0, 0, -1).Add(22 * time.Hour)

	clientRequest := signin.ListFreeSpacesRequest{
		From: startDate,
		To:   endDate,
	}
	clientResponse, err := s.signinClient.ListFreeSpaces(ctx, clientRequest)
	if err != nil {
		return resp, fmt.Errorf("calling signin client: %w", err)
	}

	for _, v := range clientResponse {
		r := ListFreeSpacesResponseItem{
			ID:       v.ID,
			DeskName: v.Name,
			SiteName: v.Zones[0].Name,
		}

		resp = append(resp, r)

	}

	return resp, nil
}
