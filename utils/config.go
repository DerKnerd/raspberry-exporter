package utils

import (
	"errors"
	"flag"
	"fmt"
	"os"

	"gopkg.in/yaml.v2"
)

const (
	defaultListenAddress = ":9549"
	defaultMetricsPath   = "/metrics"
)

type LocalConfig struct {
	Listen    ListenConfig    `yaml:"listen"`
	Raspberry RaspberryConfig `yaml:"raspberry"`
}

type ListenConfig struct {
	Address     string `yaml:"address"`
	MetricsPath string `yaml:"metricspath"`
}

type RaspberryConfig struct {
	VcGenCmd         string `yaml:"vcgencmd"`
	DisableThrottled bool   `yaml:"disable_throttled"`
	Model            string `yaml:"model"`
}

func ParseConfig() (*LocalConfig, error) {
	configFile := flag.String("config.file", "", "Path to configuration file.")
	flag.Parse()

	if *configFile == "" {
		return defaultConfig()
	}

	file, err := os.Open(*configFile)
	if err != nil {
		return nil, fmt.Errorf("can not open config file: %s", err)
	}

	config := &LocalConfig{}
	if err := yaml.NewDecoder(file).Decode(config); err != nil {
		return nil, fmt.Errorf("error decoding config file %q: %s", *configFile, err)
	}

	return config, nil
}

func getVcGenCmd() (string, error) {
	if _, err := os.Stat("/opt/vc/bin/vcgencmd"); !os.IsNotExist(err) {
		return "/opt/vc/bin/vcgencmd", nil
	} else if _, err := os.Stat("/usr/bin/vcgencmd"); !os.IsNotExist(err) {
		return "/usr/bin/vcgencmd", nil
	} else {
		return "", errors.New("could not find vcgencmd: please install the raspberry pi toolchain")
	}
}

func defaultConfig() (*LocalConfig, error) {
	vcGenCmd, err := getVcGenCmd()
	if err != nil {
		return nil, err
	}

	return &LocalConfig{
		Raspberry: RaspberryConfig{
			VcGenCmd:         vcGenCmd,
			DisableThrottled: false,
		},
		Listen: ListenConfig{
			Address:     defaultListenAddress,
			MetricsPath: defaultMetricsPath,
		},
	}, nil
}
