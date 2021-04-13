package measure

import (
	"context"

	"go.opencensus.io/stats"
)

type LatencyMeasure struct {
	measure *stats.Float64Measure
}

func NewLatencyMeasure(name, description string) *LatencyMeasure {
	return &LatencyMeasure{
		measure: stats.Float64(name, description, "ms"),
	}
}

func (m *LatencyMeasure) Record(ctx context.Context, milliseconds float64) {
	stats.Record(ctx, m.measure.M(milliseconds))
}

func (m *LatencyMeasure) Measure() stats.Measure {
	return m.measure
}
