package signin

import (
	"context"
	"fmt"
	"strings"

	"github.com/spf13/cobra"

	"github.com/cgxarrie-go/signin/config"
	"github.com/cgxarrie-go/signin/pkg/service"
	"github.com/cgxarrie-go/signin/pkg/signin"
)

var (
	idColLength   int = 10
	dateColLength int = 25
	deskColLength int = 25
	siteColLength int = 10
)

var listBookingsCmd = &cobra.Command{
	Use:     "list-bookings",
	Aliases: []string{"lb"},
	Short:   "book a desk",
	Args:    cobra.ExactArgs(1),
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

		printListBookingTitle()

		format := getListBookingColumnLineFormat()
		for _, v := range resp {
			printListBookingInfo(format, v)
		}

		return nil
	},
}

func printListBookingTitle() {
	format := getListBookingColumnTitleFormat()
	head := fmt.Sprintf(format, "ID", "Date", "Office", "Desk")
	line := strings.Repeat("-", len(head)+5)

	fmt.Println(head)
	fmt.Println(line)
}

func printListBookingInfo(format string, item service.ListBookingsResponseItem) {
	info := fmt.Sprintf(format, item.ID, item.Date.Format("2006-01-02"),
		item.SiteName, item.DeskName)

	fmt.Println(info)
}

func getListBookingColumnLineFormat() string {
	return "%" + fmt.Sprintf("%d", idColLength) + "d " +
		"| %" + fmt.Sprintf("%d", dateColLength) + "s " +
		"| %" + fmt.Sprintf("%d", deskColLength) + "s " +
		"| %" + fmt.Sprintf("%d", siteColLength) + "s "
}

func getListBookingColumnTitleFormat() string {
	return "%" + fmt.Sprintf("%d", idColLength) + "s " +
		"| %" + fmt.Sprintf("%d", dateColLength) + "s " +
		"| %" + fmt.Sprintf("%d", deskColLength) + "s " +
		"| %" + fmt.Sprintf("%d", siteColLength) + "s "
}
