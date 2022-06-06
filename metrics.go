package metrics

import "time"

func SinceInMs(start time.Time) int64 {
	return time.Since(start).Milliseconds()
}