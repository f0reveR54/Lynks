package metrics

import (
	"time"

	"github.com/prometheus/client_golang/prometheus"
)

var (
	requestCounter = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "api_requests_total",
			Help: "Total number of API requests",
		},
		[]string{"method"},
	)
	requestDuration = prometheus.NewHistogramVec(
		prometheus.HistogramOpts{
			Name:    "api_request_duration_seconds",
			Help:    "Histogram of API request durations",
			Buckets: prometheus.DefBuckets,
		},
		[]string{"method"},
	)
)

func Metrics(api string) {

	start := time.Now()
	defer func() {
		requestDuration.WithLabelValues(api).Observe(time.Since(start).Seconds())
		requestCounter.WithLabelValues(api).Inc()
	}()

}

func New() {

	prometheus.MustRegister(requestCounter)
	prometheus.MustRegister(requestDuration)
}
