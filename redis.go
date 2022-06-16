package metrics

import (
	"go.opentelemetry.io/otel/metric/global"
	"go.opentelemetry.io/otel/metric/instrument"
	"go.opentelemetry.io/otel/metric/instrument/syncfloat64"
)

func getRedisHistogramName() string {
	return getPrefix() + "." + redisHistogramName
}

var (
	meter      = global.MeterProvider().Meter(getPrefix() + ".redis")
	measure, _ = meter.SyncFloat64().Histogram(
		getRedisHistogramName(),
		instrument.WithDescription("redis duration in seconds"),
		instrument.WithUnit(Seconds),
	)
)

func GetRedisHistogram() syncfloat64.Histogram {
	return measure
}
