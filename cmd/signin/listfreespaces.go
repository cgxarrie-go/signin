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

var listFreeSpacesCmd = &cobra.Command{
	Use:     "list-free",
	Aliases: []string{"lf"},
	Short:   "book a desk",
	Args:    cobra.ExactArgs(1),
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

		printLisFreeSpacesTitle()

		format := getListFreeSpacesColumnLineFormat()
		for _, v := range resp {
			printListFreeSpacesInfo(format, v)
		}

		return nil
	},
}

func printLisFreeSpacesTitle() {
	format := getListFreeSpacesColumnTitleFormat()
	head := fmt.Sprintf(format, "ID", "Office", "Desk")
	line := strings.Repeat("-", len(head)+5)

	fmt.Println(head)
	fmt.Println(line)
}

func printListFreeSpacesInfo(format string,
	item service.ListFreeSpacesResponseItem) {
	info := fmt.Sprintf(format, item.ID,
		item.SiteName, item.DeskName)

	fmt.Println(info)
}

func getListFreeSpacesColumnLineFormat() string {
	return "%" + fmt.Sprintf("%d", idColLength) + "s " +
		"| %" + fmt.Sprintf("%d", deskColLength) + "s " +
		"| %" + fmt.Sprintf("%d", siteColLength) + "s "
}

func getListFreeSpacesColumnTitleFormat() string {
	return "%" + fmt.Sprintf("%d", idColLength) + "s " +
		"| %" + fmt.Sprintf("%d", deskColLength) + "s " +
		"| %" + fmt.Sprintf("%d", siteColLength) + "s "
}
