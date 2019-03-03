package main

import (
	"./exporter"
	"./utils"
	"flag"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"log"
	"net/http"
)

func main() {
	flag.Parse()
	utils.ParseConfig()

	exp := exporter.New()

	prometheus.MustRegister(exp)

	listenAddress := utils.Config().Listen.Address

	http.Handle(utils.Config().Listen.MetricsPath, promhttp.Handler())
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, utils.Config().Listen.MetricsPath, http.StatusMovedPermanently)
	})

	log.Printf("Starting Raspberry PI exporter on %q", listenAddress)

	if err := http.ListenAndServe(listenAddress, nil); err != nil {
		log.Fatalf("Cannot start Raspberry PI exporter: %s", err)
	} else {
		log.Printf("Started Raspberry PI exporter on %q", listenAddress)
	}
}
