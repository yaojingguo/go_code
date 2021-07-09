package main

import (
	"net/http"
	"os"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func recordMetrics() {
	go func() {
		for {
			opsProcessed.Inc()
			time.Sleep(2 * time.Second)
		}
	}()
}

var (
	opsProcessed = promauto.NewCounter(prometheus.CounterOpts{
		Name: "yao_ops_count",
		Help: "The total count of yao ops",
	})
)

func main() {
	recordMetrics()

	var port string
	if len(os.Args) <= 1 {
		port = "2112"
	} else {
		port = os.Args[1]
	}

	http.Handle("/metrics", promhttp.Handler())
	http.ListenAndServe(":"+port, nil)
}
