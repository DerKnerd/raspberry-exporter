package exporter

import (
	"../collector"
	"github.com/prometheus/client_golang/prometheus"
)

type Exporter struct {
	collectors []prometheus.Collector
}

func New() Exporter {
	var exporter = Exporter{}
	exporter.collectors = []prometheus.Collector{
		collector.NewVcGenCmdCollector(),
	}

	return exporter
}

func (exporter Exporter) Describe(channel chan<- *prometheus.Desc) {
	for _, cc := range exporter.collectors {
		cc.Describe(channel)
	}
}

func (exporter Exporter) Collect(channel chan<- prometheus.Metric) {
	for _, c := range exporter.collectors {
		c.Collect(channel)
	}
}

var _ prometheus.Collector = &Exporter{}
