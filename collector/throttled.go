package collector

import (
	"github.com/derknerd/raspberry-exporter/utils"
	"github.com/prometheus/client_golang/prometheus"
	"regexp"
	"strconv"
)

const (
	Throttled string = "get_throttled"
)

var (
	throttledRegex = regexp.MustCompile(`(throttled=)|(0x)|(\n)|(\r)`)
)

func (c *VcGenCmdCollector) getThrottledAtBit(desc *prometheus.Desc, atBit uint, device string) prometheus.Metric {
	bits, err := utils.ExecuteVcGen(c.VcGenCmd, "get_throttled", device)

	if err != nil {
		return prometheus.NewInvalidMetric(desc, err)
	}

	bits = throttledRegex.ReplaceAllString(bits, "")
	bitsUInt, err := strconv.ParseUint(bits, 16, 64)
	bitBool := 0.0
	if bitsUInt&(1<<atBit) == 0 {
		bitBool = 0
	} else {
		bitBool = 1
	}
	if err != nil {
		return prometheus.NewInvalidMetric(desc, err)
	}

	return prometheus.MustNewConstMetric(
		desc,
		prometheus.GaugeValue,
		bitBool,
	)
}
