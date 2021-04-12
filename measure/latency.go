package measure

import (
	"context"

	"go.opencensus.io/tag"

	"go.opencensus.io/stats"
)

type LatencyMeasure struct {
	measure *stats.Float64Measure
	TagKeys []tag.Key
}

func NewLatencyMeasure(name, description string, tagKeys []tag.Key) *LatencyMeasure {
	return &LatencyMeasure{
		measure: stats.Float64(name, description, "ms"),
		TagKeys: tagKeys,
	}
}

func (m *LatencyMeasure) Record(ctx context.Context, milliseconds float64) {
	stats.Record(ctx, m.measure.M(milliseconds))
}

func (m *LatencyMeasure) Measure() stats.Measure {
	return m.measure
}

func (m *LatencyMeasure) Tags() []tag.Key {
	return m.TagKeys
}
