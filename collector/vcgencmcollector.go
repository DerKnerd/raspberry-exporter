package collector

import (
	"github.com/prometheus/client_golang/prometheus"
)

var (
	prefix       = "pi_vcgencmd_"
	coreTempDesc = prometheus.NewDesc(
		prefix+"core_temp",
		"Temperature of the SoC",
		nil,
		nil,
	)
	coreVoltageDesc = prometheus.NewDesc(
		prefix+"core_voltage",
		"Voltage of the CPU",
		nil,
		nil,
	)
	sdramControllerVoltageDesc = prometheus.NewDesc(
		prefix+"sdram_controller_voltage",
		"Voltage of the SDRAM controller",
		nil,
		nil,
	)
	sdramInputOutputVoltageDesc = prometheus.NewDesc(
		prefix+"sdram_input_output_voltage",
		"Voltage of the SDRAM IO",
		nil,
		nil,
	)
	sdramPhysicalVoltageDesc = prometheus.NewDesc(
		prefix+"sdram_physical_voltage",
		"Voltage of the physical SDRAM",
		nil,
		nil,
	)
	coreClockDesc = prometheus.NewDesc(
		prefix+"gpu_clock",
		"Clock speed of the GPU in Hz",
		nil,
		nil,
	)
	armClockDesc = prometheus.NewDesc(
		prefix+"cpu_clock",
		"Clock speed of the ARM CPU in Hz",
		nil,
		nil,
	)
	emmcClockDesc = prometheus.NewDesc(
		prefix+"emmc_clock",
		"Clock speed of the SD card in Hz",
		nil,
		nil,
	)
	cpuMemoryDesc = prometheus.NewDesc(
		prefix+"cpu_memory",
		"Memory for the CPU in Megabytes",
		nil,
		nil,
	)
	gpuMemoryDesc = prometheus.NewDesc(
		prefix+"gpu_memory",
		"Memory for the GPU in Megabytes",
		nil,
		nil,
	)
	throttledUnderVoltageDetectedDesc = prometheus.NewDesc(
		prefix+"throttled_under_voltage_detected",
		"Under-voltage detected",
		nil,
		nil,
	)
	throttledArmFrequencyCappedDesc = prometheus.NewDesc(
		prefix+"throttled_arm_freq_capped",
		"Arm frequency capped",
		nil,
		nil,
	)
	throttledCurrentlyThrottled = prometheus.NewDesc(
		prefix+"throttled_currently_throttled",
		"Currently throttled",
		nil,
		nil,
	)
	throttledSoftTemperatureLimitActive = prometheus.NewDesc(
		prefix+"throttled_soft_temperature_limit_active",
		"Soft temperature limit active",
		nil,
		nil,
	)
)

type VcGenCmdCollector struct {
	VcGenCmd string
}

func (c *VcGenCmdCollector) Describe(channel chan<- *prometheus.Desc) {
	channel <- coreTempDesc
	channel <- coreVoltageDesc
	channel <- sdramControllerVoltageDesc
	channel <- sdramInputOutputVoltageDesc
	channel <- sdramPhysicalVoltageDesc
	channel <- coreClockDesc
	channel <- armClockDesc
	channel <- emmcClockDesc
	channel <- cpuMemoryDesc
	channel <- gpuMemoryDesc
	channel <- throttledUnderVoltageDetectedDesc
	channel <- throttledArmFrequencyCappedDesc
	channel <- throttledCurrentlyThrottled
	channel <- throttledSoftTemperatureLimitActive
}

func (c *VcGenCmdCollector) Collect(channel chan<- prometheus.Metric) {
	channel <- c.getCoreTemp()
	channel <- c.getVoltage(coreVoltageDesc, CoreVoltage)
	channel <- c.getVoltage(sdramPhysicalVoltageDesc, SdramPhysicalVoltage)
	channel <- c.getVoltage(sdramInputOutputVoltageDesc, SdramInputOutputVoltage)
	channel <- c.getVoltage(sdramControllerVoltageDesc, SdramControllerVoltage)
	channel <- c.getClock(coreClockDesc, GpuClock)
	channel <- c.getClock(armClockDesc, ArmClock)
	channel <- c.getClock(emmcClockDesc, EmmcClock)
	channel <- c.getMemory(cpuMemoryDesc, CpuMemory)
	channel <- c.getMemory(gpuMemoryDesc, GpuMemory)
	channel <- c.getThrottledAtBit(throttledUnderVoltageDetectedDesc, 0, Throttled)
	channel <- c.getThrottledAtBit(throttledArmFrequencyCappedDesc, 1, Throttled)
	channel <- c.getThrottledAtBit(throttledCurrentlyThrottled, 2, Throttled)
	channel <- c.getThrottledAtBit(throttledSoftTemperatureLimitActive, 3, Throttled)
}

func NewVcGenCmdCollector(vcGenCmd string) *VcGenCmdCollector {
	return &VcGenCmdCollector{
		VcGenCmd: vcGenCmd,
	}
}
