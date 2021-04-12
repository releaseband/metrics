package measure

import (
	"context"
	"time"

	"go.opencensus.io/tag"
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

func NewMutator(key tag.Key, value string, metaData ...tag.Metadata) tag.Mutator {
	return tag.Insert(key, value, metaData...)
}

func ContextFactory(mutator ...tag.Mutator) func(ctx context.Context) (context.Context, error) {
	return func(ctx context.Context) (context.Context, error) {
		return tag.New(ctx, mutator...)
	}
}
