package metrics

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

// ServerMetrics is struct for the metrics that server supports today
type ServerMetrics struct {
	ProcessedRequests *prometheus.CounterVec
}

// MetricsExporter is exporter used to export the metrics to prometheus
type MetricsExporter struct {
	Metrics  ServerMetrics // metrics that server supports
	Port     int64         // port on which exporter will be running
	Endpoint string        // endpoint which prometheus will call to get scrap metrics
}

// createMetricsObject creates type of metrics for the server
func (s *ServerMetrics) createMetricsObject() {
	s.ProcessedRequests = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "total_request",
			Help: "Number of requests processed by server",
		},
		[]string{"processed"},
	)
}

// StartExporter starts the metrics exporter
func (e *MetricsExporter) StartExporter() {
	router := mux.NewRouter()
	router.Path(e.Endpoint).Handler(promhttp.Handler())
	log.Printf("Starting metrics exporter on port %d", e.Port)
	err := http.ListenAndServe(":" + fmt.Sprintf("%d", e.Port), router)
	log.Fatal(err)
}

// NewMetricsExporter returns the exporter object with the user defined cfgd
func NewMetricsExporter(port int64, endpoint string) MetricsExporter {
	metricsObject := NewServerMetrics()
	exporter := MetricsExporter{Port: port}
	exporter.Metrics = metricsObject
	exporter.Endpoint = endpoint
	return exporter
}

// NewServerMetrics creates server metrics with metrics object that server supports
func NewServerMetrics() ServerMetrics {
	reqMetrics := ServerMetrics{}
	reqMetrics.createMetricsObject()
	prometheus.Register(reqMetrics.ProcessedRequests)
	return reqMetrics
}