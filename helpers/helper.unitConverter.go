package helpers

import "math"

const (
	SECONDS = "second"
	MINUTES = "minutes"
	HOURS   = "hours"
	DAYS    = "days"
	WEEKS   = "weeks"
	MONTHS  = "months"
)

var (
	seconds int = 1
	minutes int = 60
	hours   int = 3600
	days    int = 86400
	weeks   int = 604800
	months  int = 2630000
)

func ConvertToSeconds(interval int, intervalType string) int {
	var n int

	switch intervalType {
	case SECONDS:
		n = int(math.Ceil(math.Abs(float64((interval * seconds)))))

	case MINUTES:
		n = int(math.Ceil(math.Abs(float64((interval * minutes)))))

	case HOURS:
		n = int(math.Ceil(math.Abs(float64((interval * hours)))))

	case DAYS:
		n = int(math.Ceil(math.Abs(float64((interval * days)))))

	case WEEKS:
		n = int(math.Ceil(math.Abs(float64((interval * weeks)))))

	case MONTHS:
		n = int(math.Ceil(math.Abs(float64((interval * months)))))

	default:
		return n
	}

	return n
}
