package exporter

import (
	"github.com/prometheus/client_golang/prometheus"
)

type Exporter struct {
	collectors []prometheus.Collector
}

func New(collectors ...prometheus.Collector) Exporter {
	return Exporter{
		collectors: collectors,
	}
}

func (exporter Exporter) Describe(channel chan<- *prometheus.Desc) {
	for _, c := range exporter.collectors {
		c.Describe(channel)
	}
}

func (exporter Exporter) Collect(channel chan<- prometheus.Metric) {
	for _, c := range exporter.collectors {
		c.Collect(channel)
	}
}
