package views

import (
	"go.opencensus.io/stats"
	"go.opencensus.io/stats/view"
	"go.opencensus.io/tag"
)

type Measure interface {
	Measure() stats.Measure
}

func MakeLatencyView(name, description string, measure Measure, tags []tag.Key) *view.View {
	return &view.View{
		Name:        name,
		Description: description,
		TagKeys:     tags,
		Measure:     measure.Measure(),
		Aggregation: view.Distribution(0.2, 2, 4, 8, 16, 32, 64, 128, 256, 512, 1024,
			2048, 4096, 8192, 16384, 32768),
	}
}

func MakeCounterView(name, description string, measure Measure, tags []tag.Key) *view.View {
	return &view.View{
		Name:        name,
		Description: description,
		Measure:     measure.Measure(),
		Aggregation: view.Count(),
		TagKeys:     tags,
	}
}

func RegisterViews(views ...*view.View) error {
	return view.Register(views...)
}
