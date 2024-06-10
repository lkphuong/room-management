package utils

import (
	"fmt"
	"time"
)

func CalculateTime(start time.Time) string {
	now := time.Now().UTC().Add(7 * time.Hour)

	diff := now.Sub(start)

	hours := int(diff.Hours())
	minutes := int(diff.Minutes()) % 60
	seconds := int(diff.Seconds()) % 60

	hoursFormatted := fmt.Sprintf("%02d", hours)
	minutesFormatted := fmt.Sprintf("%02d", minutes)
	secondsFormatted := fmt.Sprintf("%02d", seconds)

	formatted := fmt.Sprintf("%s:%s:%s", hoursFormatted, minutesFormatted, secondsFormatted)

	return formatted
}
