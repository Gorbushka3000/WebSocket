package main

import (
	"log"
	"net/http"
	"ws/internal/handlers"
)

func main() {
	mux := routes()

	log.Println("Starting channel listenner")
	go handlers.ListenToWsChannel()

	log.Println("Start server")

	_ = http.ListenAndServe(":8080", mux)
}
