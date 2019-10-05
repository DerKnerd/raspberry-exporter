package collector

import (
	"regexp"
	"strconv"

	"github.com/derknerd/raspberry-exporter/utils"
	"github.com/prometheus/client_golang/prometheus"
)

const (
	CoreVoltage             string = "core"
	SdramControllerVoltage  string = "sdram_c"
	SdramInputOutputVoltage string = "sdram_i"
	SdramPhysicalVoltage    string = "sdram_p"
)

func (c *VcGenCmdCollector) getVoltage(desc *prometheus.Desc, device string) prometheus.Metric {
	voltage, err := utils.ExecuteVcGen(c.VcGenCmd, "measure_volts", device)

	if err != nil {
		return prometheus.NewInvalidMetric(desc, err)
	}

	voltage = regexp.MustCompile(`(volt=)|(V)|(\n)|(\r)`).ReplaceAllString(voltage, "")

	voltageFloat, err := strconv.ParseFloat(voltage, 64)

	if err != nil {
		return prometheus.NewInvalidMetric(desc, err)
	}

	return prometheus.MustNewConstMetric(
		desc,
		prometheus.GaugeValue,
		voltageFloat,
	)
}
