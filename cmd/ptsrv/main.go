package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/seeruk/prom-tester/internal"
)

var port int

func main() {
	flag.IntVar(&port, "port", 8080, "Port to run the HTTP server on")
	flag.Parse()

	addr := fmt.Sprintf("0.0.0.0:%d", port)
	log.Printf("Starting server on %s...\n", addr)

	resolver := internal.NewResolver()
	recorder := resolver.ResolveMetricsRecorder()
	server := resolver.ResolveHTTPServer(addr)

	recorder.Start()

	err := server.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}
