package metrics

import "os"

const (
	prefixEnvKey       = "REDIS_METRIC_PREFIX"
	redisHistogramName = "redis_duration_seconds"
)

func GetRedisHistogramName() string {
	prefix := os.Getenv(prefixEnvKey)
	if prefix == "" {
		return redisHistogramName
	}

	return prefix + "." + redisHistogramName
}
