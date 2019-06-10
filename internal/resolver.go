package internal

import (
	"net/http"

	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/seeruk/prom-tester/internal/metrics"
)

// Resolver ...
type Resolver struct {
}

// NewResolver returns a new Resolver instance.
func NewResolver() *Resolver {
	return &Resolver{}
}

// ResolveHTTPServer ...
func (r *Resolver) ResolveHTTPServer(addr string) *http.Server {
	mux := http.NewServeMux()
	mux.Handle("/metrics", promhttp.Handler())

	return &http.Server{
		Addr:    addr,
		Handler: mux,
	}
}

// ResolveMetricsRecorder ...
func (r *Resolver) ResolveMetricsRecorder() *metrics.Recorder {
	return metrics.NewRecorder()
}
