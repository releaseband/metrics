package exporter

import (
	"fmt"
	"net/http"
	"strconv"

	"contrib.go.opencensus.io/exporter/prometheus"
)

type PrometheusConfigs struct {
	Options     prometheus.Options
	MetricsPath string
	Port        int
	errHandler  func(err error)
}

func runServer(exporter *prometheus.Exporter, config PrometheusConfigs) {
	mux := http.NewServeMux()
	mux.Handle(config.MetricsPath, exporter)

	if err := http.ListenAndServe(":"+strconv.Itoa(config.Port), mux); err != nil {
		config.errHandler(err)
	}
}

func RunPrometheusExporter(config PrometheusConfigs) error {
	e, err := prometheus.NewExporter(config.Options)
	if err != nil {
		return fmt.Errorf("prometheus new Exported failed: %w", err)
	}

	go runServer(e, config)
	return nil
}
