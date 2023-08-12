package utils

import (
	"io/ioutil"
	"log"
	"os"
	"strconv"

	"gopkg.in/yaml.v2"
)

// Global variables initialised in the main package
var (
	ServerURL        string
	ServerName       string
	ServerPort       int
	ServerMaxThreads int
	ServerQueueSize  int
	ServerTokenRate  int
	ServerTokenLimit int
)

// PrometheusConfig is used for exporter
type PrometheusConfig struct {
	Global struct {
		ScrapeInterval     string `yaml:"scrape_interval"`
		EvaluationInterval string `yaml:"evaluation_interval"`
	} `yaml:"global"`
	ScrapeConfigs []struct {
		JobName       string `yaml:"job_name"`
		StaticConfigs []struct {
			Targets []string `yaml:"targets"`
		} `yaml:"static_configs"`
		MetricsPath string `yaml:"metrics_path,omitempty"`
	} `yaml:"scrape_configs"`
	FilePath         string
}

// GetConfig reads the cfg yml file 
func(p *PrometheusConfig) GetConfig() {
	if len(p.FilePath) == 0 {
		log.Print("No file present for prometheus cfg to read")
		return
	}
	yfile, err := ioutil.ReadFile(p.FilePath)
	if err != nil {
		 log.Fatal(err)
	}
	err2 := yaml.Unmarshal(yfile, p)
	if err2 != nil {
		 log.Fatal(err2)
	}
}

// GetExporterMetricsEndpoint gets the metrics endpoint for the given jobname
func(p *PrometheusConfig) GetExporterMetricsEndpoint(jobName string) string {
	for _, job := range p.ScrapeConfigs {
		if job.JobName == jobName {
			return job.MetricsPath
		}
	}
	log.Printf("No job %s found in the yml file", jobName)
	return ""
}

// GetExporterURL gets the exporter url from the yml file for the given jobname
func(p *PrometheusConfig) GetExporterURL(jobName string) string {
	for _, job := range p.ScrapeConfigs {
		if job.JobName == jobName {
			if len(job.StaticConfigs) > 0 && len(job.StaticConfigs[0].Targets) > 0 {
				return job.StaticConfigs[0].Targets[0]
			}
		}
	}
	log.Printf("No job %s found in the yml file", jobName)
	return ""
}

// SetServerConfig reads the env var and sets the global variable for the server to use 
func SetServerConfig() {
	ServerURL = os.Getenv("SERVER_URL")
	ServerName = os.Getenv("SERVER_NAME")
	ServerPort, _ = strconv.Atoi(os.Getenv("SERVER_PORT"))
	ServerMaxThreads, _ = strconv.Atoi(os.Getenv("SERVER_WORKERS"))
	ServerQueueSize, _ = strconv.Atoi(os.Getenv("SERVER_QUEUE_SIZE"))
	ServerTokenRate, _ = strconv.Atoi(os.Getenv("SERVER_TOKEN_RATE"))
	ServerTokenLimit, _ = strconv.Atoi(os.Getenv("SERVER_TOKEN_LIMIT"))

	log.Printf("Workers: %d, Queue size: %d, Token rate: %d, Token limit: %d", ServerMaxThreads,ServerQueueSize, ServerTokenRate, ServerTokenLimit)
}
