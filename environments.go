package metrics

import "os"

const (
	defaultPrefix      = "RB"
	prefixEnvKey       = "REDIS_METRIC_PREFIX"
	redisHistogramName = "redis_duration_seconds"
)

func getPrefix() string {
	prefix := os.Getenv(prefixEnvKey)
	if prefix == "" {
		return defaultPrefix
	}

	return prefix
}
