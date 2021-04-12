package main

import (
	"context"
	"fmt"
	"log"
	"strconv"
	"time"

	"contrib.go.opencensus.io/exporter/prometheus"
	"github.com/releaseband/metrics/exporter"

	"github.com/releaseband/metrics/opencensus/views"

	"github.com/releaseband/metrics/measure"
	"go.opencensus.io/tag"
)

const (
	path      = "/metrics"
	namespace = "my_app"
)

var (
	counter       *measure.CounterMeasure
	latency       *measure.LatencyMeasure
	entityKey     = tag.MustNewKey("entity")
	workerNameKey = tag.MustNewKey("name")
)

func worker(i int64) {
	workerName := strconv.FormatInt(i, 10) + "sec"

	fmt.Println("worker " + workerName + " started")

	mutators := []tag.Mutator{
		measure.NewMutator(entityKey, "worker"),
		measure.NewMutator(workerNameKey, workerName),
	}

	ctxFactory := measure.ContextFactory(mutators...)

	for {
		ctx, err := ctxFactory(context.Background())
		if err != nil {
			log.Fatal(err)
		}

		counter.IncrementCounter(ctx)

		start := measure.Start()
		time.Sleep(time.Duration(i) * time.Second)
		latency.Record(ctx, measure.End(start))
	}
}

func initViews() error {
	latencyTagKeys := []tag.Key{entityKey, workerNameKey}
	counterTagKeys := []tag.Key{entityKey, workerNameKey}

	latency = measure.NewLatencyMeasure("latency", "The latency measure in milliseconds", latencyTagKeys)
	counter = measure.NewCounterMeasure("counter", "The counter measure", counterTagKeys)

	latencyView := views.MakeLatencyView("worker #1", "latency by worker #1", latency)
	counterView := views.MakeCounterView("worker #2", "how many times the worker was called", counter)

	return views.RegisterViews(latencyView, counterView)
}

func startWorkers() {
	for i := int64(1); i < 10; i++ {
		go worker(i)
	}
}

func runServer(port int) error {
	opt := prometheus.Options{
		Namespace: namespace,
	}

	config := exporter.NewPrometheusConfigs(opt, path, port, func(err error) {
		log.Fatal(err)
	})

	return exporter.RunPrometheusExporter(config)
}

func main() {
	if err := initViews(); err != nil {
		log.Fatal(err)
	}

	startWorkers()

	if err := runServer(9999); err != nil {
		log.Fatal(err)
	}

	time.Sleep(10 * time.Minute)
}
