package opencensus

import (
	"context"
	"fmt"
	"time"

	"go.opencensus.io/stats/view"

	"go.opencensus.io/tag"

	"go.opencensus.io/stats"
)

var (
	latency     = stats.Float64("latency", "The latency in milliseconds", "ms")
	latencyHttp = stats.Float64("http latency", "Delaying HTTP requests", "ms")
	counter     = stats.Int64("counter", "counter", "1")

	keyMethod   = tag.MustNewKey("method")
	keyEntity   = tag.MustNewKey("entity")
	keyHttpCode = tag.MustNewKey("http_code")
	keyURL      = tag.MustNewKey("url")

	initialized bool
	errHandler  func(err error)
)

type OpencensusMetrics struct{}

func NewOpencensusMetrics(errorHandler func(err error)) (OpencensusMetrics, error) {
	m := OpencensusMetrics{}

	if err := view.Register(views...); err != nil {
		return m, fmt.Errorf("register views failed: %w", err)
	}

	initialized = true
	errHandler = errorHandler

	return m, nil
}

func sinceInMilliseconds(st time.Time) float64 {
	return float64(time.Since(st).Nanoseconds()) / 1e6
}

func measure(callback func()) float64 {
	start := time.Now()
	callback()
	return sinceInMilliseconds(start)
}

func MeasureLatency(ctx context.Context, entity, method string, callback func()) {
	if !initialized {
		callback()
		return
	}

	ms := measure(callback)
	ctx, err := makeLatencyCtx(ctx, entity, method)
	if err != nil {
		errHandler(fmt.Errorf("makeLatencyCtx failed: %w", err))
		return
	}

	stats.Record(ctx, latency.M(ms))
}

func MeasureHttpLatency(ctx context.Context, entity, method, url string, callback func()) {
	if !initialized {
		callback()
		return
	}

	ms := measure(callback)

	ctx, err := makeReqCtx(ctx, entity, method, url)
	if err != nil {
		errHandler(fmt.Errorf("makeReqCtx failed: %w", err))
		return
	}

	stats.Record(ctx, latencyHttp.M(ms))
}

func CommitHttpCode(ctx context.Context, url string, code int) {
	if !initialized {
		return
	}

	ctx, err := makeHttpCodeCtx(ctx, url, code)
	if err != nil {
		errHandler(fmt.Errorf("makeHttpCodeCtx failed: %w", err))
		return
	}

	stats.Record(ctx, counter.M(1))
}

func (m OpencensusMetrics) MeasureLatency(ctx context.Context, entity, method string, callback func()) {
	MeasureLatency(ctx, entity, method, callback)
}

func (m OpencensusMetrics) MeasureHttpLatency(ctx context.Context, entity, method, url string, callback func()) {
	MeasureHttpLatency(ctx, entity, method, url, callback)
}

func (m OpencensusMetrics) CommitHttpCode(ctx context.Context, url string, code int) {
	CommitHttpCode(ctx, url, code)
}
