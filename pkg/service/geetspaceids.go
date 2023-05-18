package service

import (
	"context"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/cgxarrie-go/signin/config"
	"github.com/cgxarrie-go/signin/pkg/signin"
)

// GetSpaceIDs fill spaceIds in the config
func (s service) GetSpaceIDs(ctx context.Context) (err error) {

	startDate := time.Date(2020, 1, 1, 23, 59, 59, 999, time.Now().Location())
	endDate := time.Date(2020, 1, 2, 0, 0, 0, 0, time.Now().Location())

	clientRequest := signin.ListFreeSpacesRequest{
		From: startDate,
		To:   endDate,
	}
	clientResponse, err := s.signinClient.ListFreeSpaces(ctx, clientRequest)
	if err != nil {
		return fmt.Errorf("calling signin client: %w", err)
	}

	for _, v := range clientResponse {

		nameParts := strings.Split(v.Name, " ")
		number, err := strconv.Atoi(nameParts[1])
		if err != nil {
			continue
		}

		if _, ok := config.Instance().Desks[number]; ok {
			continue
		}

		config.Instance().Desks[number] = v.ID
	}

	config.Instance().Save()

	return nil
}
