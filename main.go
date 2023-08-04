package main

import (
	"log"

	"github.com/Avish34/tcp-server/server"
)


func main() {
	log.Println("Starting the server")
	serverObject := &server.Server{Port: 8080}
	serverObject.Init()
	log.Println("Server started")
}