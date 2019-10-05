package main

import (
	"log"
	"net/http"

	"github.com/derknerd/raspberry-exporter/collector"
	"github.com/derknerd/raspberry-exporter/exporter"
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

	c := collector.NewVcGenCmdCollector(config.Raspberry.VcGenCmd)
	exp := exporter.New(c)

	prometheus.MustRegister(exp)

	listenAddress := config.Listen.Address

	http.Handle(config.Listen.MetricsPath, promhttp.Handler())
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, utils.Config().Listen.MetricsPath, http.StatusMovedPermanently)
	})

	log.Printf("Starting Raspberry PI exporter on %q", listenAddress)

	if err := http.ListenAndServe(listenAddress, nil); err != nil {
		log.Fatalf("Cannot start Raspberry PI exporter: %s", err)
	}
}
