package utils

import (
	"fmt"
	"time"
)

func FormatDate(date string) string {
	if result, err := time.Parse(time.RFC3339, date); err == nil {
		return fmt.Sprintf("%02d/%02d/%04d", result.Month(), result.Day(), result.Year())
	} else {
		return date
	}
}
