package main

import (
	"log"
	"net/http"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/collectors"
	"github.com/prometheus/client_golang/prometheus/promhttp"

	"github.com/andreschaffer/slis-slos-examples/simple-service/api/handlers"
	"github.com/andreschaffer/slis-slos-examples/simple-service/api/middleware"
)

func main() {
	registry := prometheus.NewRegistry()
	registerDefaultCollectors(registry)

	metrics := middleware.NewMetrics(registry, nil)
	registerAdminRoutes(registry, metrics)
	registerAppRoutes(registry, metrics)

	log.Fatalln(http.ListenAndServe(":8888", nil))
}

func registerDefaultCollectors(registry *prometheus.Registry) {
	registry.MustRegister(
		collectors.NewGoCollector(),
		collectors.NewProcessCollector(collectors.ProcessCollectorOpts{}),
	)
}

func registerAdminRoutes(registry *prometheus.Registry, metrics middleware.Metrics) {
	metricsRoute := "/metrics"
	metricsHandler := metrics.WrapHandler(metricsRoute, promhttp.HandlerFor(registry, promhttp.HandlerOpts{}))
	http.Handle(metricsRoute, metricsHandler)
}

func registerAppRoutes(registry *prometheus.Registry, metrics middleware.Metrics) {
	pingRoute := "/ping"
	pingHandler := metrics.WrapHandler(pingRoute, handlers.Ping())
	http.Handle(pingRoute, pingHandler)
}
