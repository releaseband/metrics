package metrics

import "go.opencensus.io/stats/view"

func DeliveryServicesAggregation() *view.Aggregation {
	// Latency in buckets:
	// [>=0ms, >=25ms, >=50ms, >=75ms, >=100ms, >=200ms, >=400ms, >=600ms, >=800ms, >=1s, >=2s, >=4s, >=6s]
	return view.Distribution(0, 25, 50, 75, 100, 200, 400, 600, 800, 1000, 2000, 4000, 6000)
}

func DefaultAggregation() *view.Aggregation {
	// Latency in buckets:
	// [>=0ms, >=0.2ms, >=2ms, >=4ms, >=8ms, >=16ms, >=32ms, >=64ms, >=128ms, >=256s, >=512s, >=1s, >=2s]
	return view.Distribution(0, 0.2, 2, 4, 8, 16, 32, 64, 128, 256, 512, 1024, 2048)
}
