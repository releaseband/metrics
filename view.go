package metrics

import (
	"go.opencensus.io/stats"
	"go.opencensus.io/stats/view"
	"go.opencensus.io/tag"
)

func NewLatencyView(name, desc string, measure stats.Measure, aggregation *view.Aggregation, tags ...tag.Key) *view.View {
	return &view.View{
		Name:        name,
		Description: desc,
		TagKeys:     tags,
		Measure:     measure,
		Aggregation: aggregation,
	}
}

func NewDeliveryServicesLatencyView(name, desc string, measure stats.Measure, tags ...tag.Key) *view.View {
	return NewLatencyView(name, desc, measure, DeliveryServicesAggregation(), tags...)
}

func NewDefaultLatencyView(name, desc string, measure stats.Measure, tags ...tag.Key) *view.View {
	return NewLatencyView(name, desc, measure, DefaultAggregation(), tags...)
}

func NewCounterView(name, description string, measure stats.Measure, tags ...tag.Key) *view.View {
	return &view.View{
		Name:        name,
		Description: description,
		Measure:     measure,
		Aggregation: view.Count(),
		TagKeys:     tags,
	}
}

func Register(w ...*view.View) error {
	return view.Register(w...)
}
