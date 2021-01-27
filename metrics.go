package metrics

import "context"

type Metrics interface {
	MeasureLatency(ctx context.Context, entity, method string, callback func())
}
