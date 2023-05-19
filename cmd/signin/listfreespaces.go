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

var listFreeSpacesCmd = &cobra.Command{
	Use:     "list-free",
	Aliases: []string{"lf"},
	Short:   "lists all the fre spaces in a given date",
	Args:    cobra.ExactArgs(1),
	Example: fmt.Sprintf("signin list-free <Date in YYYYMMDD>\n" +
		"signin list-free 20230901\n" +
		"signin lf 20230901"),
	RunE: func(cmd *cobra.Command, args []string) error {

		ctx := context.Background()

		client := signin.NewClient(config.Instance().Bearer)
		svc := service.New(client)

		req := service.ListFreeSpacesRequest{
			Date: args[0],
		}

		resp, err := svc.ListFreeSpaces(ctx, req)
		if err != nil {
			return fmt.Errorf("calling service ListFreeSpaces: %w", err)
		}

		desks := make([]ui.Desk, len(resp))
		for i, r := range resp {
			desk := ui.Desk{
				ID:       r.ID,
				Name:     r.DeskName,
				ZoneName: r.ZoneName,
			}
			desks[i] = desk
		}

		ui.Instance().PrintDesks(desks)

		return nil
	},
}
