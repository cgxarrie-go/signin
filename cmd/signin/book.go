package signin

import (
	"context"
	"fmt"

	"github.com/spf13/cobra"

	"github.com/cgxarrie-go/signin/config"
	"github.com/cgxarrie-go/signin/pkg/service"
	"github.com/cgxarrie-go/signin/pkg/signin"
	"github.com/cgxarrie-go/signin/ui"
)

var bookCmd = &cobra.Command{
	Use:     "book",
	Aliases: []string{"b"},
	Short:   "book a desk",
	Example: fmt.Sprintf("signin book <DeskNumber> <Date YYYYMMDD> <Date YYYYMMDD> ...\n" +
		"signin book 59 20230524 20230525 ...\n" +
		"signin b 59 20230524 20230525 ..."),
	RunE: func(cmd *cobra.Command, args []string) error {

		ctx := context.Background()

		client := signin.NewClient(config.Instance().Bearer)
		svc := service.New(client)

		req := service.BookSpaceRequest{
			DeskNumber: args[0],
			Dates:      args[1:],
		}

		resp, err := svc.BookSpace(ctx, req)
		if err != nil {
			return fmt.Errorf("calling service BookSpace: %w", err)
		}

		bookings := []ui.Booking{}
		for _, item := range resp {
			booking := ui.Booking{
				ID: item.ID,
				Desk: ui.Desk{
					ID:       item.DeskID,
					Name:     item.DeskName,
					ZoneName: item.ZoneName,
				},
				Date: item.Date,
			}
			bookings = append(bookings, booking)
		}

		ui.Instance().PrintBookings(bookings)

		return nil
	},
}
