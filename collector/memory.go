package collector

import (
	"regexp"
	"strconv"

	"github.com/derknerd/raspberry-exporter/utils"
	"github.com/prometheus/client_golang/prometheus"
)

const (
	CpuMemory string = "arm"
	GpuMemory string = "gpu"
)

func (c *VcGenCmdCollector) getMemory(desc *prometheus.Desc, device string) prometheus.Metric {
	memory, err := utils.ExecuteVcGen(c.VcGenCmd, "get_mem", device)

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
