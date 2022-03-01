package metrics

import (
	"context"
	"go.opencensus.io/stats"
	"go.opencensus.io/tag"
)

func NewCounter(name, description string) *stats.Int64Measure {
	return stats.Int64(name, description, "1")
}

func AddCount(ctx context.Context, m *stats.Int64Measure, v int64, mutators ...tag.Mutator) (context.Context, error) {
	newCtx, err := tag.New(ctx, mutators...)
	if err != nil {
		return nil, err
	}

	stats.Record(newCtx, m.M(v))

	return newCtx, nil
}
