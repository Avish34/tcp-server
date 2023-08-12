package main

import (
	"log"
	"strconv"
	"strings"

	"github.com/Avish34/tcp-server/metrics"
	"github.com/Avish34/tcp-server/server"
	"github.com/Avish34/tcp-server/utils"
	. "github.com/Avish34/tcp-server/utils"
)



func main() {
	// Get the env vars
	utils.SetServerConfig()
	log.Printf("Starting the server on %s:%d", ServerURL, ServerPort)

	// Set the exporter configs and create metrics exporter object
	filePath := "./configs/prometheus.yml"
	metricsCfg := utils.PrometheusConfig{FilePath: filePath}
	metricsCfg.GetConfig()
	metricsEndpoint := metricsCfg.GetExporterMetricsEndpoint(ServerName)
	metricsPort, _ := strconv.Atoi(strings.Split(metricsCfg.GetExporterURL(ServerName), ":")[1])
	exporter := metrics.NewMetricsExporter(int64(metricsPort), metricsEndpoint)

	// Server options
	opts := server.ServerOpts{
		MaxThreads: ServerMaxThreads,
		QueueSize: ServerQueueSize,
		Rate: int64(ServerTokenRate),
		Tokens: int64(ServerTokenLimit),
	}
	// Create server object
	serverObject := &server.Server{Port: ServerPort, URL: ServerURL, Opts: opts, Metrics: exporter.Metrics}

	// Start metrics exporter and server
	go exporter.StartExporter()
	serverObject.FireUpTheServer()
	log.Println("Server and metrics exporter started")

}