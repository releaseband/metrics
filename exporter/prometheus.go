package exporter

import (
	"fmt"
	"net/http"
	"strconv"

	"contrib.go.opencensus.io/exporter/prometheus"
)

type PrometheusConfigs struct {
	options     prometheus.Options
	metricsPath string
	port        int
	errHandler  func(err error)
}

func NewPrometheusConfigs(
	options prometheus.Options,
	metricsPath string,
	port int,
	errorHandler func(err error),
) PrometheusConfigs {
	return PrometheusConfigs{
		options:     options,
		metricsPath: metricsPath,
		port:        port,
		errHandler:  errorHandler,
	}
}

func runServer(exporter *prometheus.Exporter, config PrometheusConfigs) {
	mux := http.NewServeMux()
	mux.Handle(config.metricsPath, exporter)

	if err := http.ListenAndServe(":"+strconv.Itoa(config.port), mux); err != nil {
		config.errHandler(err)
	}
}

func RunPrometheusExporter(config PrometheusConfigs) error {
	e, err := prometheus.NewExporter(config.options)
	if err != nil {
		return fmt.Errorf("prometheus new Exported failed: %w", err)
	}

	go runServer(e, config)
	return nil
}
