package measure

import (
	"context"

	"go.opencensus.io/stats"
)

type CounterMeasure struct {
	measure *stats.Int64Measure
}

func NewCounterMeasure(name, description string) *CounterMeasure {
	return &CounterMeasure{
		measure: stats.Int64(name, description, "1"),
	}
}

func (m *CounterMeasure) IncrementCounter(ctx context.Context) {
	stats.Record(ctx, m.measure.M(1))
}

func (m *CounterMeasure) Measure() stats.Measure {
	return m.measure
}
