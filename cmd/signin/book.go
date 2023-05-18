package signin

import (
	"context"
	"fmt"

	"github.com/spf13/cobra"

	"github.com/cgxarrie-go/signin/config"
	"github.com/cgxarrie-go/signin/pkg/service"
	"github.com/cgxarrie-go/signin/pkg/signin"
)

var bookCmd = &cobra.Command{
	Use:     "book",
	Aliases: []string{"b"},
	Short:   "book a desk",
	Args:    cobra.ExactArgs(2),
	RunE: func(cmd *cobra.Command, args []string) error {

		ctx := context.Background()

		client := signin.NewClient(config.Instance().Bearer)
		svc := service.New(client)

		req := service.BookSpaceRequest{
			DeskNumber: args[0],
			Date:       args[1],
		}

		resp, err := svc.BookSpace(ctx, req)
		if err != nil {
			return fmt.Errorf("calling service BookSpace: %w", err)
		}

		fmt.Printf("%+v", resp)

		return nil
	},
}
