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
	Short:   "cancel desk reservations for a date",
	Args:    cobra.ExactArgs(1),
	Example: fmt.Sprintf("signin cancel <date>\n" +
		"signin cancel 20231008\n" +
		"signin c 202308"),
	RunE: func(cmd *cobra.Command, args []string) error {

		ctx := context.Background()

		client := signin.NewClient(config.Instance().Bearer)
		svc := service.New(client)

		req := service.CancelBookingRequest{
			Date: args[0],
		}

		err := svc.CancelBooking(ctx, req)
		if err != nil {
			return fmt.Errorf("calling service BookSpace: %w", err)
		}

		fmt.Println("---------------------")
		fmt.Println("Booking cancelled")
		fmt.Println("---------------------")

		return nil
	},
}
