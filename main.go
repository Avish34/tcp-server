package main

import (
	"log"

	"github.com/Avish34/tcp-server/server"
)


func main() {
	log.Println("Starting the server")
	serverObject := &server.Server{Port: 8080, URL: "0.0.0.0", MaxThreads: 2}
	serverObject.FireUpTheServer()
	log.Println("Server started")
}