package main

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"net/http"
	"time"
)

func main() {
	age := prometheus.NewCounterVec(prometheus.CounterOpts{
		Name: "age",
		Help: "age by second",
	}, []string{"age"})
	age.WithLabelValues("value")

	register := prometheus.NewRegistry()
	register.MustRegister(age)
	go func() {
		for {
			age.WithLabelValues("value").Inc()
			time.Sleep(1 * time.Second)
		}
	}()
	http.Handle("/metrics", promhttp.HandlerFor(register, promhttp.HandlerOpts{Registry: register}))
	err := http.ListenAndServe(":1888", nil)
	if err != nil {
		panic("error listen")
		return
	}
}
