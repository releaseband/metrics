package measure

import (
	"context"

	"go.opencensus.io/tag"

	"go.opencensus.io/stats"
)

type CounterMeasure struct {
	measure *stats.Int64Measure
	tagKeys []tag.Key
}

func NewCounterMeasure(name, description string, tagKeys []tag.Key) *CounterMeasure {
	return &CounterMeasure{
		tagKeys: tagKeys,
		measure: stats.Int64(name, description, "1"),
	}
}

func (m *CounterMeasure) IncrementCounter(ctx context.Context) {
	stats.Record(ctx, m.measure.M(1))
}

func (m *CounterMeasure) Measure() stats.Measure {
	return m.measure
}

func (m *CounterMeasure) Tags() []tag.Key {
	return m.tagKeys
}
