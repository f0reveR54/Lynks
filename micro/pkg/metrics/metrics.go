package metrics

import (
	"github.com/prometheus/client_golang/prometheus"
)

var (
	RequestCounter = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "api_requests_total",
			Help: "Total number of API requests",
		},
		[]string{"method"},
	)
	RequestDuration = prometheus.NewHistogramVec(
		prometheus.HistogramOpts{
			Name:    "api_request_duration_seconds",
			Help:    "Histogram of API request durations",
			Buckets: prometheus.DefBuckets,
		},
		[]string{"method"},
	)
)

func New() {

	prometheus.MustRegister(RequestCounter)
	prometheus.MustRegister(RequestDuration)

}
