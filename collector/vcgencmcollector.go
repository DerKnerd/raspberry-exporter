package collector

import (
	"io/ioutil"
	"log"
	"strings"

	"github.com/derknerd/raspberry-exporter/utils"
	"github.com/prometheus/client_golang/prometheus"
)

func getModel(cfg utils.RaspberryConfig) string {
	if cfg.Model != "" {
		return cfg.Model
	}
	data, err := ioutil.ReadFile("/proc/cpuinfo")
	if err != nil {
		log.Printf("failed to read /proc/cpuinfo: %s", err)
		return "unknown"
	}

	for _, line := range strings.Split(string(data), "\n") {
		if !strings.HasPrefix(line, "Model\t") {
			continue
		}
		parts := strings.SplitN(line, ":", 2)
		if len(parts) == 2 {
			return strings.TrimSpace(parts[1])
		}
	}
	return "unknown"
}

var (
	prefix                              = "pi_vcgencmd_"
	coreTempDesc                        *prometheus.Desc
	coreVoltageDesc                     *prometheus.Desc
	sdramControllerVoltageDesc          *prometheus.Desc
	sdramInputOutputVoltageDesc         *prometheus.Desc
	sdramPhysicalVoltageDesc            *prometheus.Desc
	coreClockDesc                       *prometheus.Desc
	armClockDesc                        *prometheus.Desc
	emmcClockDesc                       *prometheus.Desc
	cpuMemoryDesc                       *prometheus.Desc
	gpuMemoryDesc                       *prometheus.Desc
	throttledUnderVoltageDetectedDesc   *prometheus.Desc
	throttledArmFrequencyCappedDesc     *prometheus.Desc
	throttledCurrentlyThrottled         *prometheus.Desc
	throttledSoftTemperatureLimitActive *prometheus.Desc
)

func initDescriptions(model string) {

	constLabels := prometheus.Labels{
		"model": model,
	}

	coreTempDesc = prometheus.NewDesc(
		prefix+"core_temp",
		"Temperature of the SoC",
		nil,
		constLabels,
	)
	coreVoltageDesc = prometheus.NewDesc(
		prefix+"core_voltage",
		"Voltage of the CPU",
		nil,
		constLabels,
	)
	sdramControllerVoltageDesc = prometheus.NewDesc(
		prefix+"sdram_controller_voltage",
		"Voltage of the SDRAM controller",
		nil,
		constLabels,
	)
	sdramInputOutputVoltageDesc = prometheus.NewDesc(
		prefix+"sdram_input_output_voltage",
		"Voltage of the SDRAM IO",
		nil,
		constLabels,
	)
	sdramPhysicalVoltageDesc = prometheus.NewDesc(
		prefix+"sdram_physical_voltage",
		"Voltage of the physical SDRAM",
		nil,
		constLabels,
	)
	coreClockDesc = prometheus.NewDesc(
		prefix+"gpu_clock",
		"Clock speed of the GPU in Hz",
		nil,
		constLabels,
	)
	armClockDesc = prometheus.NewDesc(
		prefix+"cpu_clock",
		"Clock speed of the ARM CPU in Hz",
		nil,
		constLabels,
	)
	emmcClockDesc = prometheus.NewDesc(
		prefix+"emmc_clock",
		"Clock speed of the SD card in Hz",
		nil,
		constLabels,
	)
	cpuMemoryDesc = prometheus.NewDesc(
		prefix+"cpu_memory",
		"Memory for the CPU in Megabytes",
		nil,
		constLabels,
	)
	gpuMemoryDesc = prometheus.NewDesc(
		prefix+"gpu_memory",
		"Memory for the GPU in Megabytes",
		nil,
		constLabels,
	)
	throttledUnderVoltageDetectedDesc = prometheus.NewDesc(
		prefix+"throttled_under_voltage_detected",
		"Under-voltage detected",
		nil,
		constLabels,
	)
	throttledArmFrequencyCappedDesc = prometheus.NewDesc(
		prefix+"throttled_arm_freq_capped",
		"Arm frequency capped",
		nil,
		constLabels,
	)
	throttledCurrentlyThrottled = prometheus.NewDesc(
		prefix+"throttled_currently_throttled",
		"Currently throttled",
		nil,
		constLabels,
	)
	throttledSoftTemperatureLimitActive = prometheus.NewDesc(
		prefix+"throttled_soft_temperature_limit_active",
		"Soft temperature limit active",
		nil,
		constLabels,
	)
}

type VcGenCmdCollector struct {
	VcGenCmd         string
	DisableThrottled bool
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
	if !c.DisableThrottled {
		channel <- throttledUnderVoltageDetectedDesc
		channel <- throttledArmFrequencyCappedDesc
		channel <- throttledCurrentlyThrottled
		channel <- throttledSoftTemperatureLimitActive
	}
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
	if !c.DisableThrottled {
		channel <- c.getThrottledAtBit(throttledUnderVoltageDetectedDesc, 0, Throttled)
		channel <- c.getThrottledAtBit(throttledArmFrequencyCappedDesc, 1, Throttled)
		channel <- c.getThrottledAtBit(throttledCurrentlyThrottled, 2, Throttled)
		channel <- c.getThrottledAtBit(throttledSoftTemperatureLimitActive, 3, Throttled)
	}
}

func NewVcGenCmdCollector(cfg utils.RaspberryConfig) *VcGenCmdCollector {
	model := getModel(cfg)
	initDescriptions(model)
	return &VcGenCmdCollector{
		VcGenCmd:         cfg.VcGenCmd,
		DisableThrottled: cfg.DisableThrottled,
	}
}
