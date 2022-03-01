package main

import (
	"context"
	"fmt"
	"github.com/releaseband/metrics"
	"log"
	"strconv"
	"time"

	"contrib.go.opencensus.io/exporter/prometheus"
	"github.com/releaseband/metrics/exporter"

	"go.opencensus.io/tag"
)

const (
	path      = "/metrics"
	namespace = "my_app"
)

var (
	command = tag.MustNewKey("command")
	id      = tag.MustNewKey("id")
	method  = tag.MustNewKey("method")
)

var (
	latency     = metrics.NewLatency("latency_name", "latency description")
	httpLatency = metrics.NewLatency("http_latency", "http latency")
	counter     = metrics.NewCounter("iterations", "iterations")
)

func a(ctx context.Context, i int) {
	ctx, record, err := metrics.Start(context.Background(), latency,
		tag.Insert(command, "a"),
		tag.Insert(id, strconv.Itoa(i)),
	)

	if err != nil {
		panic(err)
	}

	time.Sleep(time.Duration(i) * time.Millisecond)

	record()
}

func httpSend(ctx context.Context, i int) {
	ctx, record, err := metrics.Start(ctx, httpLatency,
		tag.Insert(method, "POST"),
		tag.Insert(id, strconv.Itoa(i)),
	)

	if err != nil {
		panic(err)
	}

	if i%2 == 0 {
		time.Sleep(200 * time.Millisecond)
	} else {
		time.Sleep(50 * time.Millisecond)
	}

	record()
}

func c(ctx context.Context, i int) {
	ctx, record, err := metrics.Start(ctx, latency,
		tag.Insert(command, "c"),
		tag.Insert(id, strconv.Itoa(i)),
	)

	if err != nil {
		panic(err)
	}

	time.Sleep(time.Duration(i) * time.Millisecond)

	httpSend(ctx, i)

	record()
}

func b(ctx context.Context, i int) {
	newCtx, record, err := metrics.Start(ctx, latency,
		tag.Insert(command, "b"),
		tag.Insert(id, strconv.Itoa(i)),
	)

	if err != nil {
		panic(err)
	}

	c(ctx, i) //

	time.Sleep(time.Duration(i) * time.Millisecond)

	record()

	_ = newCtx
}

func worker(i int) {
	fmt.Println("worker ", i, " started")

	for {
		ctx := context.Background()

		a(ctx, i)

		b(ctx, i)

		if _, err := metrics.AddCount(ctx, counter, 1, tag.Insert(id, strconv.Itoa(i))); err != nil {
			panic(err)
		}
	}
}

func initViews() error {
	latencyView := metrics.NewDefaultLatencyView("latency_view",
		"latency view description", latency, command, id)
	httpLatencyView := metrics.NewDeliveryServicesLatencyView("http_latency", "http service latency",
		httpLatency, method, id)
	counterView := metrics.NewCounterView("counter_view",
		"counter view description", counter, id)

	return metrics.Register(latencyView, counterView, httpLatencyView)
}

func startWorkers() {
	for i := 0; i < 64; i++ {
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
