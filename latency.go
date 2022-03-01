package metrics

import (
	"context"
	"go.opencensus.io/stats"
	"go.opencensus.io/tag"
	"time"
)

func convertToMilliseconds(t int64) float64 {
	return float64(t) / 1e6
}

func end(start time.Time) float64 {
	return convertToMilliseconds(time.Since(start).Nanoseconds())
}

func recordLatency(ctx context.Context, m *stats.Float64Measure, start time.Time) {
	stats.Record(ctx, m.M(end(start)))
}

func Start(ctx context.Context, m *stats.Float64Measure, mutators ...tag.Mutator) (context.Context, func(), error) {
	newCtx, err := tag.New(ctx, mutators...)
	if err != nil {
		return nil, nil, err
	}

	start := time.Now()

	end := func() {
		recordLatency(newCtx, m, start)
	}

	return newCtx, end, nil
}

func NewLatency(name, description string) *stats.Float64Measure {
	return stats.Float64(name, description, "ms")
}
