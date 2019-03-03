package collector

import (
	"../utils"
	"github.com/prometheus/client_golang/prometheus"
	"regexp"
	"strconv"
)

const (
	CpuMemory string = "arm"
	GpuMemory string = "gpu"
)

func getMemory(desc *prometheus.Desc, device string) prometheus.Metric {
	memory, err := utils.ExecuteVcGen("get_mem", device)

	if err != nil {
		return prometheus.NewInvalidMetric(desc, err)
	}

	memory = regexp.MustCompile(`(gpu=)|(arm=)|(M)|(\n)|(\r)`).ReplaceAllString(memory, "")
	memoryFloat, err := strconv.ParseFloat(memory, 64)

	if err != nil {
		return prometheus.NewInvalidMetric(desc, err)
	}

	return prometheus.MustNewConstMetric(
		desc,
		prometheus.GaugeValue,
		memoryFloat,
	)
}
