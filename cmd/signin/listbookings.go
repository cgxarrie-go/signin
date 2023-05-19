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

var listBookingsCmd = &cobra.Command{
	Use:     "list-bookings",
	Aliases: []string{"lb"},
	Short:   "list the active user bookings until a given date",
	Args:    cobra.ExactArgs(1),
	Example: fmt.Sprintf("signin list-bookings <Date in YYYYMMDD>\n" +
		"signin list-bookings 20230901\n" +
		"signin lb 20230901"),
	RunE: func(cmd *cobra.Command, args []string) error {
		ctx := context.Background()

		client := signin.NewClient(config.Instance().Bearer)
		svc := service.New(client)

		req := service.ListBookingsquest{
			EndDate: args[0],
		}

		resp, err := svc.ListBookings(ctx, req)
		if err != nil {
			return fmt.Errorf("calling service ListBookings: %w", err)
		}

		bookings := make([]ui.Booking, len(resp))
		for i, r := range resp {
			booking := ui.Booking{
				ID: r.ID,
				Desk: ui.Desk{
					ID:       r.DeskID,
					Name:     r.DeskName,
					ZoneName: r.ZoneName,
				},
				Date: r.Date,
			}

			bookings[i] = booking
		}

		ui.Instance().PrintBookings(bookings)
		return nil
	},
}
