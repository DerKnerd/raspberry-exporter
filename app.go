package main

import (
	"log"
	"net/http"

	"github.com/derknerd/raspberry-exporter/collector"
	"github.com/derknerd/raspberry-exporter/utils"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func main() {
	log.SetFlags(0)

	config, err := utils.ParseConfig()
	if err != nil {
		log.Fatalf("Error loading configuration: %s", err)
	}

	c := collector.NewVcGenCmdCollector(config.Raspberry)
	prometheus.MustRegister(c)

	listenAddress := config.Listen.Address

	http.Handle(config.Listen.MetricsPath, promhttp.Handler())
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, config.Listen.MetricsPath, http.StatusFound)
	})

	log.Printf("Starting Raspberry PI exporter on %q", listenAddress)

	if err := http.ListenAndServe(listenAddress, nil); err != nil {
		log.Fatalf("Cannot start Raspberry PI exporter: %s", err)
	}
}
