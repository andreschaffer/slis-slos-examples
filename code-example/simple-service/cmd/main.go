package main

import (
	"log"
	"net/http"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/collectors"
	"github.com/prometheus/client_golang/prometheus/promhttp"

	"github.com/andreschaffer/slis-slos-examples/simple-service/internal/api/handlers"
	"github.com/andreschaffer/slis-slos-examples/simple-service/internal/api/middleware"
)

func main() {
	registry := prometheus.NewRegistry()
	registerDefaultCollectors(registry)

	metrics := middleware.NewMetrics(registry)
	adminMux := adminMux(registry, metrics)
	appMux := appMux(registry, metrics)

	go func() {
		log.Fatalln(http.ListenAndServe(":8889", adminMux))
	}()
	log.Fatalln(http.ListenAndServe(":8888", appMux))
}

func registerDefaultCollectors(registry *prometheus.Registry) {
	registry.MustRegister(
		collectors.NewGoCollector(),
		collectors.NewProcessCollector(collectors.ProcessCollectorOpts{}),
	)
}

func adminMux(registry *prometheus.Registry, metrics middleware.Metrics) *http.ServeMux {
	serverMux := http.NewServeMux()
	metricsRoute := "/metrics"
	metricsHandler := metrics.WrapHandler(metricsRoute, promhttp.HandlerFor(registry, promhttp.HandlerOpts{}))
	serverMux.Handle(metricsRoute, metricsHandler)
	return serverMux

}

func appMux(registry *prometheus.Registry, metrics middleware.Metrics) *http.ServeMux {
	serverMux := http.NewServeMux()
	pingRoute := "/ping"
	pingHandler := metrics.WrapHandler(pingRoute, handlers.Ping())
	serverMux.Handle(pingRoute, pingHandler)
	return serverMux
}
