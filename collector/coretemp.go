package collector

import (
	"io/ioutil"
	"regexp"
	"strconv"
	"strings"

	"github.com/derknerd/raspberry-exporter/utils"
	"github.com/prometheus/client_golang/prometheus"
)

func (c *VcGenCmdCollector) getCoreTemp() prometheus.Metric {
	coreTempFromSys, err := ioutil.ReadFile("/sys/class/thermal/thermal_zone0/temp")
	var coreTemp string
	method := "file"

	if err != nil {
		if coreTemp, err = utils.ExecuteVcGen(c.VcGenCmd, "measure_temp"); err != nil {
			return prometheus.NewInvalidMetric(coreTempDesc, err)
		} else {
			coreTemp = regexp.MustCompile(`(temp=)|('C)|(\n)|(\r)`).ReplaceAllString(coreTemp, "")
			method = "vcgen"
		}
	} else {
		coreTemp = string(coreTempFromSys)
	}

	coreTemp = strings.TrimSuffix(coreTemp, "\n")
	coreTempFloat, err := strconv.ParseFloat(coreTemp, 64)

	if err != nil {
		return prometheus.NewInvalidMetric(coreTempDesc, err)
	}

	if method == "file" {
		coreTempFloat /= 1000
	}

	return prometheus.MustNewConstMetric(
		coreTempDesc,
		prometheus.GaugeValue,
		coreTempFloat,
	)
}
