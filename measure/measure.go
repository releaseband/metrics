package measure

import (
	"time"
)

func Start() time.Time {
	return time.Now()
}

func convertToMilliseconds(t int64) float64 {
	return float64(t) / 1e6
}

// time in milliseconds
func End(start time.Time) float64 {
	return convertToMilliseconds(time.Since(start).Nanoseconds())
}
