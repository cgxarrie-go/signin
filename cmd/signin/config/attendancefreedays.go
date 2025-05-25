package config

import (
	"fmt"
	"strconv"

	"github.com/spf13/cobra"

	appcfg "github.com/cgxarrie-go/signin/config"
	"github.com/cgxarrie-go/signin/pkg/util"
)

// azurePATCmd represents the azurePAT command
var attendanceFreeDaysCmd = &cobra.Command{
	Use:     "attendance-free-days",
	Aliases: []string{"afd"},
	Short:   "set attendace free days",
	Args:    cobra.ExactArgs(3),
	Example: fmt.Sprintf("signin config attendance-free-days <start-date> <number-of-days> <reason>\n" +
		"signin config afd <start-date> <number-of-days> <reason>"),
	RunE: func(cmd *cobra.Command, args []string) error {
		err := runAttendanceFreeDaysCmd(args)
		return err
	},
}

func init() {
	ConfigCmd.AddCommand(attendanceFreeDaysCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// azurePATCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// azurePATCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func runAttendanceFreeDaysCmd(args []string) error {
	cfg := appcfg.Instance()
	cfg.Load()

	startDate, err := util.DateFromString(args[0])
	if err != nil {
		return fmt.Errorf("invalid start date: %s", err.Error())
	}
	numberOfDays, err := strconv.Atoi(args[1])
	if err != nil {
		return fmt.Errorf("invalid number of days: %s", err.Error())
	}
	if numberOfDays < 1 {
		return fmt.Errorf("number of days should be greater than 0")
	}
	reason := args[2]
	if reason == "" {
		return fmt.Errorf("reason should not be empty")
	}
	// Add the free days to the config
	if cfg.AttendanceFreeDays == nil {
		cfg.AttendanceFreeDays = make(map[string]string)
	}
	for i := 0; i < numberOfDays; i++ {
		date := startDate.AddDate(0, 0, i)
		cfg.AttendanceFreeDays[date.Format("2006-01-02")] = reason
	}

	// Save the config
	err = cfg.Save()
	if err != nil {
		return err
	}
	return nil
}
