package opencensus

import (
	"go.opencensus.io/stats/view"
	"go.opencensus.io/tag"
)

var (
	latencyBounds = []float64{0.2, 2, 4, 8, 16, 32, 64, 128, 256, 512, 1024, 2048, 4096,
		8192, 16384, 32768}
	views = []*view.View{
		{
			Name:        "latency",
			Description: "The various latencies of the method",
			Measure:     latency,
			Aggregation: view.Distribution(latencyBounds...),
			TagKeys:     []tag.Key{keyEntity, keyMethod},
		},
		{
			Name:        "http requests",
			Description: "http requests latency",
			Measure:     latencyHttp,
			Aggregation: view.Distribution(latencyBounds...),
			TagKeys:     []tag.Key{keyEntity, keyURL, keyMethod},
		},
		{
			Name:        "http codes",
			Description: "http code counts",
			Measure:     counter,
			Aggregation: view.Count(),
			TagKeys:     []tag.Key{keyHttpCode, keyURL},
		},
	}
)
