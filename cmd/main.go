package main

import (
	"diagram-server/internal/bootstrap"
	"log"
	"net/http"
)

func main() {
	port := ":8080"

	bootstrap.StartUp(port)

	log.Printf(" Server is ready to handle requests on %s", port)
	if err := http.ListenAndServe(port, nil); err != nil {
		log.Fatal(err)
	}
}
