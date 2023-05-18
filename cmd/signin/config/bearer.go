package config

import (
	"github.com/spf13/cobra"

	appcfg "github.com/cgxarrie-go/signin/config"
)

// azurePATCmd represents the azurePAT command
var bearerCmd = &cobra.Command{
	Use:   "bearer",
	Short: "set the signin Bearer",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		err := runBearerCmd(args)
		return err
	},
}

func init() {
	ConfigCmd.AddCommand(bearerCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// azurePATCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// azurePATCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func runBearerCmd(args []string) error {
	cfg := appcfg.Instance()
	cfg.Load()

	cfg.Bearer = args[0]
	err := cfg.Save()
	if err != nil {
		return err
	}
	return nil
}
