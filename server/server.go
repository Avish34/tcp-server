package server

import (
	"fmt"
	"log"
	"net"

	"github.com/Avish34/tcp-server/metrics"
)

// Server for accepting tcp connections and processing the request
type Server struct {
	WorkerPool              			// Collection of threads
	Port       int          			// Port on which server should listen
	URL        string       			// URL for the server to run
	Opts       ServerOpts   			// ServerOpts will be used to customize the server
	Metrics    metrics.ServerMetrics    // Metrics will be used to export to prometheus 
	Listener   net.Listener 			// Socket listener
	reqLimiter TokenBucket  			// reqLimiter will the throttle the number of request processed
}

// ServerOpts is used by Server to customize it as per the user
type ServerOpts struct {
	Rate 	   int64 // Rate at which bucket should be filled with tokens, defined per seconds
	Tokens     int64 // Number of Tokens for rate limiter 
	MaxThreads int   // Number of threads can that be used
	QueueSize  int   // Request to be put in queue if workers are busy
}

func (s *Server) FireUpTheServer() {
	s.createListener()
	s.createThreadPool()
	s.createRateLimiter(s.Opts.Rate, s.Opts.Tokens)
	s.handleRequest()
}

// createListener creates a socket listener on the given port
func (s *Server) createListener() {
	log.Println("Creating listener")
	address := s.URL + ":" + fmt.Sprintf("%d",s.Port)
	socketObject, err := net.Listen("tcp", address)
	if err != nil {
		log.Fatalf("Server creation failed with %v", err)
	}
	s.Listener = socketObject
	
}

// createThreadPool creates a worker pool based on the threads given
func (s *Server) createThreadPool() {
	log.Println("Creating thread pool")
	s.WorkerPool = WorkerPool{MaxWorkers: s.Opts.MaxThreads, QueueSize: s.Opts.QueueSize}
	s.NewWorkerPool()
}

func (s *Server) createRateLimiter(rate, token int64) {
	log.Println("Creating rate limiter")
	s.reqLimiter = NewRateLimiter(rate, token)
}

// handleRequest handles the new connections
func (s *Server) handleRequest() {
	log.Println("Start handling requests")
	conn := 0
	for {
		client, err := s.Listener.Accept() // blocking
		conn++
		log.Printf("Client connected %d", conn)
		if err != nil {
			log.Fatalf("%v", err)
		}
		log.Println("Processing the request")
		if s.reqLimiter != (TokenBucket{}) && s.reqLimiter.isRequestAllowed() {
			s.Metrics.ProcessedRequests.WithLabelValues("processed").Inc()
			s.SubmitJob(Job{JobId: conn, Conn: client})
		} else {
			log.Println("You have reached your API limit, please wait before re-trying")
			client.Close()
		}
	}
}

// Close closes the socket listener and worker pool
func (s *Server) Close() {
	s.Listener.Close()
	s.WorkerPool.Close()
}