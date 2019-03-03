package collector

import (
	"github.com/prometheus/client_golang/prometheus"
)

type VcGenCmdCollector struct {
	CoreTemp                *prometheus.Desc
	CoreVoltage             *prometheus.Desc
	SdramControllerVoltage  *prometheus.Desc
	SdramInputOutputVoltage *prometheus.Desc
	SdramPhysicalVoltage    *prometheus.Desc
	CoreClock               *prometheus.Desc
	ArmClock                *prometheus.Desc
	EmmcClock               *prometheus.Desc
}

func (VcGenCmdCollector) Describe(channel chan<- *prometheus.Desc) {
	var collector = NewVcGenCmdCollector()

	channel <- collector.CoreTemp
	channel <- collector.CoreVoltage
	channel <- collector.SdramControllerVoltage
	channel <- collector.SdramInputOutputVoltage
	channel <- collector.SdramPhysicalVoltage
	channel <- collector.CoreClock
	channel <- collector.ArmClock
	channel <- collector.EmmcClock
}

func (VcGenCmdCollector) Collect(channel chan<- prometheus.Metric) {
	var collector = NewVcGenCmdCollector()

	channel <- getCoreTemp(*collector)
	channel <- getVoltage(collector.CoreVoltage, CoreVoltage)
	channel <- getVoltage(collector.SdramPhysicalVoltage, SdramPhysicalVoltage)
	channel <- getVoltage(collector.SdramInputOutputVoltage, SdramInputOutputVoltage)
	channel <- getVoltage(collector.SdramControllerVoltage, SdramControllerVoltage)
	channel <- getClock(collector.CoreClock, CoreClock)
	channel <- getClock(collector.ArmClock, ArmClock)
	channel <- getClock(collector.EmmcClock, EmmcClock)
}

var _ prometheus.Collector = &VcGenCmdCollector{}

func NewVcGenCmdCollector() *VcGenCmdCollector {
	const (
		subsystem = "vcgencmd"
		namespace = "pi"
	)

	return &VcGenCmdCollector{
		CoreTemp: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, subsystem, "core_temp"),
			"Temperature of the SoC",
			nil,
			nil,
		),
		CoreVoltage: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, subsystem, "core_voltage"),
			"Voltage of the CPU",
			nil,
			nil,
		),
		SdramControllerVoltage: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, subsystem, "sdram_controller_voltage"),
			"Voltage of the SDRAM controller",
			nil,
			nil,
		),
		SdramInputOutputVoltage: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, subsystem, "sdram_input_output_voltage"),
			"Voltage of the SDRAM IO",
			nil,
			nil,
		),
		SdramPhysicalVoltage: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, subsystem, "sdram_physical_voltage"),
			"Voltage of the physical SDRAM",
			nil,
			nil,
		),
		CoreClock: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, subsystem, "core_clock"),
			"Clock of the core in Hz",
			nil,
			nil,
		),
		ArmClock: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, subsystem, "arm_clock"),
			"Clock of the ARM CPU in Hz",
			nil,
			nil,
		),
		EmmcClock: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, subsystem, "emmc_clock"),
			"Clock of the external MMC in Hz",
			nil,
			nil,
		),
	}
}
