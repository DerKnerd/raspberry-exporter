package collector

import (
	"../utils"
	"github.com/prometheus/client_golang/prometheus"
	"regexp"
	"strconv"
)

const (
	CoreClock string = "core"
	EmmcClock string = "emmc"
	ArmClock  string = "arm"
)

func getClock(desc *prometheus.Desc, device string) prometheus.Metric {
	clock, err := utils.ExecuteVcGen("measure_clock", device)

	if err != nil {
		return prometheus.NewInvalidMetric(desc, err)
	}

	clock = regexp.MustCompile(`frequency\(\d*\)=|(\n)|(\r)`).ReplaceAllString(clock, "")

	clockFloat, err := strconv.ParseFloat(clock, 64)

	if err != nil {
		return prometheus.NewInvalidMetric(desc, err)
	}

	return prometheus.MustNewConstMetric(
		desc,
		prometheus.GaugeValue,
		clockFloat,
	)
}
