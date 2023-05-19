package signin

import (
	"context"
	"fmt"

	"github.com/spf13/cobra"

	"github.com/cgxarrie-go/signin/config"
	"github.com/cgxarrie-go/signin/pkg/service"
	"github.com/cgxarrie-go/signin/pkg/signin"
)

var cancelBooking = &cobra.Command{
	Use:     "cancel",
	Aliases: []string{"c"},
	Short:   "cancel a desk reservation",
	Args:    cobra.ExactArgs(1),
	Example: fmt.Sprintf("signin cancel <ReservationID>\n" +
		"signin cancel 57387ghf8\n" +
		"signin c 57387ghf8"),
	RunE: func(cmd *cobra.Command, args []string) error {

		ctx := context.Background()

		client := signin.NewClient(config.Instance().Bearer)
		svc := service.New(client)

		req := service.CancelBookingRequest{
			BookingID: args[0],
		}

		err := svc.CancelBooking(ctx, req)
		if err != nil {
			return fmt.Errorf("calling service BookSpace: %w", err)
		}

		fmt.Printf("Booking cancelled")

		return nil
	},
}
