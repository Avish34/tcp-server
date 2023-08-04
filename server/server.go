package server

import (
	"fmt"
	"log"
	"net"
	"time"
)

type Server struct {
	Port int // Port on which server should listen
	ThreadPool int // Number of threads can that be used
	URL string // URL for the server to run
	Listener net.Listener // Socket listener
}

func (s *Server) Init() {
	s.createListener()
	s.createThreadPool()
	s.handleRequest()
}

func (s *Server) createListener() {
	log.Println("Creating listener")
	address := s.URL + ":" + fmt.Sprintf("%d",s.Port)
	socketObject, err := net.Listen("tcp", address)
	if err != nil {
		log.Fatalf("Server creation failed with %v", err)
	}
	s.Listener = socketObject
}

func (s *Server) createThreadPool() {
	// TODO
	log.Println("Creating thread pool")
}

func (s *Server) handleRequest() {
	log.Println("Start handling requests")
	for {
		client, err := s.Listener.Accept() // blocking
		log.Printf("Client connected")
		if err != nil {
			log.Fatalf("%v", err)
		}
		log.Println("Processing the request")
		go s.processRequest(client)
	}
}

func (s *Server) processRequest(conn net.Conn) {
	request := make([]byte, 1024)
	conn.Read(request)
	response := []byte("HTTP/1.1 200 OK\r\n\r\n Hello world ! \r\n")
	time.Sleep(3*time.Second)
	conn.Write(response)
	conn.Close()
}