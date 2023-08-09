package metrics

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

type ServerMetrics struct {
	ProcessedRequests *prometheus.CounterVec
}

type MetricsExporter struct {
	Metrics  ServerMetrics
	Port     int64
	Endpoint string
}

func (s *ServerMetrics) createMetricsObject() {
	s.ProcessedRequests = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "total_request",
			Help: "Number of requests processed by server",
		},
		[]string{"processed"},
	)
}

func (e *MetricsExporter) StartExporter() {
	router := mux.NewRouter()
	router.Path(e.Endpoint).Handler(promhttp.Handler())

	log.Printf("Starting metrics exporter on port %d", e.Port)
	err := http.ListenAndServe(":" + fmt.Sprintf("%d", e.Port), router)
	log.Fatal(err)
}

func NewMetricsExporter(port int64, endpoint string) MetricsExporter {
	metricsObject := NewServerMetrics()
	exporter := MetricsExporter{Port: port}
	exporter.Metrics = metricsObject
	exporter.Endpoint = endpoint
	return exporter
}

func NewServerMetrics() ServerMetrics {
	reqMetrics := ServerMetrics{}
	reqMetrics.createMetricsObject()
	prometheus.Register(reqMetrics.ProcessedRequests)
	return reqMetrics
}