package server

import (
	"fmt"
	"log"
	"net"
)

// Server for accepting tcp connections and processing the request
type Server struct {
	Port       int          // Port on which server should listen
	MaxThreads int          // Number of threads can that be used
	URL        string       // URL for the server to run
	Listener   net.Listener // Socket listener
	WorkerPool              // Collection of threads
}

func (s *Server) FireUpTheServer() {
	s.createListener()
	s.createThreadPool()
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
	s.WorkerPool = WorkerPool{MaxWorkers: s.MaxThreads, QueueSize: 3}
	s.NewWorkerPool()
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
		s.SubmitJob(Job{JobId: conn, Conn: client})
	}
}

// Close closes the socket listener and worker pool
func (s *Server) Close() {
	s.Listener.Close()
	s.WorkerPool.Close()
}