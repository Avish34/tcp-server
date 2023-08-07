package main

import (
	"log"

	"github.com/Avish34/tcp-server/server"
)


func main() {
	log.Println("Starting the server")
	opts := server.ServerOpts{
		MaxThreads: 2,
		QueueSize: 3,
		Rate: 1,
		Tokens: 1,
	}
	serverObject := &server.Server{Port: 8080, URL: "0.0.0.0", Opts: opts}
	serverObject.FireUpTheServer()
	log.Println("Server started")
}