package signin

import (
	"context"
	"fmt"
	"strconv"

	"github.com/spf13/cobra"

	"github.com/cgxarrie-go/signin/config"
	"github.com/cgxarrie-go/signin/pkg/service"
	"github.com/cgxarrie-go/signin/pkg/signin"
	"github.com/cgxarrie-go/signin/ui"
)

var attendanceCmd = &cobra.Command{
	Use:     "attendance",
	Aliases: []string{"a"},
	Short:   "attendance per week",
	Example: fmt.Sprintf("signin attendance <NumerOfWeeks>\n" +
		"signin attendance 6\n" +
		"signin b 6"),
	RunE: func(cmd *cobra.Command, args []string) error {

		ctx := context.Background()

		client := signin.NewClient(config.Instance().Bearer)
		svc := service.New(client)

		weeks, err := parseAttendanceArgs(args)
		if err != nil {
			return err
		}

		req := service.AttendanceRequest{
			NumberOfWeeks: weeks,
		}

		resp, err := svc.Attendance(ctx, req)
		if err != nil {
			return fmt.Errorf("calling service BookSpace: %w", err)
		}

		attendanceWeeks := map[int]ui.AttendanceWeek{}
		for _, item := range resp.Items {
			attendanceWeekKey := item.RelativeWeek + len(resp.Items) - 1
			week := ui.AttendanceWeek{
				RelativeWeek:     item.RelativeWeek,
				WeekStartDate:    item.WeekStartDate,
				WeekEndDate:      item.WeekEndDate,
				Week:             item.Week,
				NumberOfBookings: item.NumberOfBookings,
				NumberOfVisits:   item.NumberOfVisits,
			}
			attendanceWeeks[attendanceWeekKey] = week
		}

		attendanceSummary := ui.AttendanceSummary{
			AverageVisitsPerWeek:   resp.Summary.AverageVisitsPerWeek,
			AverageBookingsPerWeek: resp.Summary.AverageBookingsPerWeek,
		}
		attendance := ui.Attendance{
			Weeks:   attendanceWeeks,
			Summary: attendanceSummary,
		}

		ui.Instance().PrintAttendance(attendance)

		return nil
	},
}

func parseAttendanceArgs(args []string) (weeks int, err error) {

	weeks, err = strconv.Atoi(args[0])
	if err == nil { // arg is a number
		if len(args) != 1 {
			return weeks, fmt.Errorf("invalid number of arguments")
		}

		return weeks, nil
	}

	return 6, nil
}
