package utils

import (
	"fmt"
	"time"

	"github.com/lkphuong/room-management/configs/hardcode"
)

func CalculateTime(start string) string {

	if start == "" {
		return ""
	}

	t, err := time.Parse(hardcode.DATETIME_LAYOUT, start)
	if err != nil {
		return ""
	}

	now := time.Now().UTC().Add(7 * time.Hour)

	diff := now.Sub(t)

	hours := int(diff.Hours())
	minutes := int(diff.Minutes()) % 60
	seconds := int(diff.Seconds()) % 60

	hoursFormatted := fmt.Sprintf("%02d", hours)
	minutesFormatted := fmt.Sprintf("%02d", minutes)
	secondsFormatted := fmt.Sprintf("%02d", seconds)

	formatted := fmt.Sprintf("%s:%s:%s", hoursFormatted, minutesFormatted, secondsFormatted)

	return formatted
}

func FormatDateString(date string) (string, error) {
	t, err := time.Parse(time.RFC3339, date)
	if err != nil {
		return "", err
	}
	return t.Format(hardcode.DATETIME_LAYOUT), nil
}
