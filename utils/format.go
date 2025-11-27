package utils

import (
	"math"
	"time"
)

func FormatDate(date time.Time) string {
	return date.Format("2006-01-02")
}

func NumberPrecision(number float64) float64 {
	return math.Round(number*100) / 100
}
