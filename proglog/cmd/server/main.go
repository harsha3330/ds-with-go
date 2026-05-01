package main

import (
	"flag"
	"log"

	"github.com/harsha3330/proglog/internal/server"
)

func main() {
	port := flag.String("port", "8080", "Application Port")
	flag.Parse()

	if *port == "" {
		log.Fatal("port cannot be empty")
	}

	addr := ":" + *port

	srv := server.NewHTTPServer(addr)

	log.Printf("Starting server on %s\n", addr)
	log.Fatal(srv.ListenAndServe())
}
