package config

import (
	"encoding/json"
	"fmt"

	"github.com/spf13/cobra"

	appcfg "github.com/cgxarrie-go/signin/config"
)

// ConfigCmd represents the Config command
var ConfigCmd = &cobra.Command{
	Use:   "config",
	Short: "display config",
	Long:  `display config`,
	RunE: func(cmd *cobra.Command, args []string) error {
		return runConfigCmd(cmd, args)
	},
}

func runConfigCmd(cmd *cobra.Command, args []string) error {

	cfg := appcfg.Instance()
	cfg.Load()
	b, err := json.Marshal(cfg)
	if err != nil {
		return err
	}
	fmt.Println(string(b))
	return nil
}
