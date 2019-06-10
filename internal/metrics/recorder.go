package metrics

import (
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

var (
	processedOpsTotal = promauto.NewCounter(prometheus.CounterOpts{
		Name: "promtest_processed_ops_total",
		Help: "The total number of processed operations",
	})
)

// Recorder ...
type Recorder struct{}

// NewRecorder returns a new Recorder instance.
func NewRecorder() *Recorder {
	return &Recorder{}
}

// Start ...
func (r *Recorder) Start() {
	go func() {
		ticker := time.NewTicker(time.Second)

		for {
			select {
			case <-ticker.C:
				r.handleProcessedOps()
			}
		}
	}()
}

// handleProcessedOps ...
func (r *Recorder) handleProcessedOps() {
	processedOpsTotal.Inc()
}
