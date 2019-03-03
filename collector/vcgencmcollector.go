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
	CpuMemory               *prometheus.Desc
	GpuMemory               *prometheus.Desc
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
	channel <- collector.CpuMemory
	channel <- collector.GpuMemory
}

func (VcGenCmdCollector) Collect(channel chan<- prometheus.Metric) {
	var collector = NewVcGenCmdCollector()

	channel <- getCoreTemp(*collector)
	channel <- getVoltage(collector.CoreVoltage, CoreVoltage)
	channel <- getVoltage(collector.SdramPhysicalVoltage, SdramPhysicalVoltage)
	channel <- getVoltage(collector.SdramInputOutputVoltage, SdramInputOutputVoltage)
	channel <- getVoltage(collector.SdramControllerVoltage, SdramControllerVoltage)
	channel <- getClock(collector.CoreClock, GpuClock)
	channel <- getClock(collector.ArmClock, ArmClock)
	channel <- getClock(collector.EmmcClock, EmmcClock)
	channel <- getMemory(collector.CpuMemory, CpuMemory)
	channel <- getMemory(collector.GpuMemory, GpuMemory)
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
			prometheus.BuildFQName(namespace, subsystem, "gpu_clock"),
			"Clock speed of the GPU in Hz",
			nil,
			nil,
		),
		ArmClock: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, subsystem, "cpu_clock"),
			"Clock speed of the ARM CPU in Hz",
			nil,
			nil,
		),
		EmmcClock: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, subsystem, "emmc_clock"),
			"Clock speed of the SD card in Hz",
			nil,
			nil,
		),
		CpuMemory: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, subsystem, "cpu_memory"),
			"Memory for the CPU in Megabytes",
			nil,
			nil,
		),
		GpuMemory: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, subsystem, "gpu_memory"),
			"Memory for the GPU in Megabytes",
			nil,
			nil,
		),
	}
}
