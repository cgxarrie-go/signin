package signin

import (
	"context"
	"fmt"
	"strconv"
	"time"

	"github.com/spf13/cobra"

	"github.com/cgxarrie-go/signin/config"
	"github.com/cgxarrie-go/signin/pkg/service"
	"github.com/cgxarrie-go/signin/pkg/signin"
	"github.com/cgxarrie-go/signin/ui"
)

var listFree bool

var attendanceCmd = &cobra.Command{
	Use:     "attendance",
	Aliases: []string{"a"},
	Short:   "attendance per week",
	Example: fmt.Sprintf("signin attendance <NumerOfWeeks>\n" +
		"signin attendance 6\n" +
		"signin a 6"),
	RunE: func(cmd *cobra.Command, args []string) error {

		if listFree {
			return execAttendanceFreeDaysList()
		}

		return execAttendanceReport(args)

	},
}

func init() {
	attendanceCmd.Flags().BoolVarP(&listFree, "listfree", "f", false, "List free attendance days")
}

func execAttendanceReport(args []string) error {
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
			RelativeWeek:        item.RelativeWeek,
			WeekStartDate:       item.WeekStartDate,
			WeekEndDate:         item.WeekEndDate,
			Week:                item.Week,
			WorkingDays:         item.WorkingDays,
			Bookings:            item.Bookings,
			Visits:              item.Visits,
			VisitsPerWorkingDay: 100 * item.VisitsPerWorkingDay,
		}

		attendanceWeeks[attendanceWeekKey] = week
	}

	attendanceSummary := ui.AttendanceSummary{
		WorkingDays:         resp.Summary.WorkingDays,
		Visits:              resp.Summary.Visits,
		AvgOfficeTimeToDay:  100 * resp.Summary.AvgOfficeTimeToDay,
		AvgOfficeTimeToWeek: 100 * resp.Summary.AvgOfficeTimeToWeek,
	}

	attendance := ui.Attendance{
		Weeks:   attendanceWeeks,
		Summary: attendanceSummary,
	}

	ui.Instance().PrintAttendance(attendance)

	return nil
}

func execAttendanceFreeDaysList() error {

	days := ui.AttendanceFreeDays{}

	cfg := config.Instance().AttendanceFreeDays

	for i := -100; i <= 0; i++ {
		day := time.Now().AddDate(0, 0, i)
		if item, ok := cfg[day.Format("2006-01-02")]; ok {
			days.FreeDays = append(days.FreeDays, ui.AttendanceFreeDay{
				Date:   day,
				Reason: item,
			})
		}
	}

	ui.Instance().PrintAttendanceFreeDays(days)
	return nil
}

func parseAttendanceArgs(args []string) (weeks int, err error) {

	if len(args) == 0 {
		return 12, nil
	}

	weeks, err = strconv.Atoi(args[0])
	if err != nil {
		return 0, fmt.Errorf("parsing weeks: %w", err)
	}

	return weeks, nil
}
