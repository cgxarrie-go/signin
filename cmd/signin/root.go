package signin

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"

	"github.com/cgxarrie-go/signin/cmd/signin/config"
	cfg "github.com/cgxarrie-go/signin/config"
)

var rootCmd = &cobra.Command{
	Use:   "signin",
	Short: "signin - a simple CLI to book places at the office",
	Long:  "signin is a simple CLI to book places at the office using sign in API",
	Run: func(cmd *cobra.Command, args []string) {

	},
}

func init() {
	rootCmd.AddCommand(config.ConfigCmd)
	rootCmd.AddCommand(bookCmd)
	rootCmd.AddCommand(cancelBooking)
	rootCmd.AddCommand(listFreeSpacesCmd)
	rootCmd.AddCommand(listBookingsCmd)
}

// Execute starts the CLI execution
func Execute() {

	cfg.Instance().Load()

	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "Error processing command :  %s\n", err)
		os.Exit(1)

	}
}
