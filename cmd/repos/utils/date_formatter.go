package utils

import (
	"time"

	"github.com/goodsign/monday"
)

func FormatDate(date string) string {
	if result, err := time.Parse(time.RFC3339, date); err == nil {
		return monday.Format(result, "02 de January de 2006", monday.LocaleEsES)
	} else {
		return date
	}
}
