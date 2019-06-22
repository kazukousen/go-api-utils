package metrics

import (
	"github.com/prometheus/client_golang/prometheus"
)

// HTTPRequestLatency ...
var HTTPRequestLatency = prometheus.NewHistogramVec(prometheus.HistogramOpts{
	Namespace: "app",
	Name:      "http_request_latency",
	Help:      "the latency of the HTTP requests in milliseconds",
	Buckets:   []float64{10, 100, 500, 1000, 5000},
}, []string{"method", "code"})

func init() {
	prometheus.MustRegister(HTTPRequestLatency)
}
