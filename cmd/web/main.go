package main

import (
	"log"
	"net/http"
)

func main() {
	mux := routes()
	log.Println("Start server")
	_ = http.ListenAndServe(":8080", mux)
}
