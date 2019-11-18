package util

import (
	"fmt"
	"time"
)

const (
	panelHeight = 7  // start from Sunday
	panelWidth  = 52 // Ignoring current week
)

// DateToPanel will return X, Y (0,0 is top-left)
func DateToPanel(currDate, inputDate time.Time) (int, int) {
	// Get Y, based on day
	y := inputDate.Weekday()

	// Get X, based on week
	currSunday := currDate.AddDate(0, 0, -1*int(currDate.Weekday()))
	inputSunday := inputDate.AddDate(0, 0, -1*int(inputDate.Weekday()))

	diff := currSunday.Sub(inputSunday)
	diffInWeek := (int(diff.Hours()) / 24) / panelHeight

	fmt.Printf("%v", diffInWeek)
	fmt.Println()

	x := panelWidth - diffInWeek
	return x, int(y)
}

// PanelToDate will return the actual date
func PanelToDate(currDate time.Time, x, y int) time.Time {
	currSunday := currDate.AddDate(0, 0, -1*int(currDate.Weekday()))

	diffInWeek := panelWidth - x
	actualDate := currSunday.AddDate(0, 0, y)
	actualDate = actualDate.AddDate(0, 0, -1*diffInWeek*7)

	return actualDate
}
