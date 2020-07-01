package middleware

import (
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

// From https://github.com/DanielHeckrath/gin-prometheus/blob/master/gin_prometheus.go
func computeApproximateRequestSize(r *http.Request) int {
	s := 0
	if r.URL != nil {
		s = len(r.URL.Path)
	}

	s += len(r.Method)
	s += len(r.Proto)
	for name, values := range r.Header {
		s += len(name)
		for _, value := range values {
			s += len(value)
		}
	}
	s += len(r.Host)

	// N.B. r.Form and r.MultipartForm are assumed to be included in r.URL.

	if r.ContentLength != -1 {
		s += int(r.ContentLength)
	}
	return s
}

func Prometheus() gin.HandlerFunc {
	requestCountMetrics := promauto.NewCounterVec(prometheus.CounterOpts{
		Subsystem: "http",
		Name:      "requests_total",
		Help:      "How many HTTP requests processed, partitioned by status code and HTTP method.",
	}, []string{"status", "method", "path"})
	requestDurationMetrics := promauto.NewHistogramVec(prometheus.HistogramOpts{
		Subsystem: "http",
		Name:      "request_duration_seconds",
		Help:      "The HTTP request latencies in seconds.",
	}, []string{"status", "method", "path"})
	requestSizeMetrics := promauto.NewSummary(prometheus.SummaryOpts{
		Subsystem: "http",
		Name:      "request_size_bytes",
		Help:      "The HTTP response sizes in bytes.",
	})
	responseSizeMetrics := promauto.NewSummary(prometheus.SummaryOpts{
		Subsystem: "http",
		Name:      "response_size_bytes",
		Help:      "The HTTP request sizes in bytes.",
	})

	return func(ctx *gin.Context) {
		start := time.Now()
		ctx.Next()
		dur := float64(time.Since(start)) / float64(time.Second)
		status := strconv.Itoa(ctx.Writer.Status())
		method := ctx.Request.Method
		url := ctx.FullPath()
		reqSz := computeApproximateRequestSize(ctx.Request)
		respSz := float64(ctx.Writer.Size())

		requestCountMetrics.WithLabelValues(status, method, url).Inc()
		requestDurationMetrics.WithLabelValues(status, method, url).Observe(dur)
		requestSizeMetrics.Observe(float64(reqSz))
		responseSizeMetrics.Observe(respSz)
	}
}

func PrometheusHandler() gin.HandlerFunc {
	return gin.WrapH(promhttp.Handler())
}
