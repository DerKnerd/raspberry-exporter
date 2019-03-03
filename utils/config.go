package utils

import (
	"flag"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
)

const (
	defaultListenAddress = ":9549"
	defaultMetricsPath   = "/metrics"
)

type localConfig struct {
	Listen    listen    `yaml:"listen"`
	Raspberry raspberry `yaml:"raspberry"`
}

type listen struct {
	Address     string `yaml:"address"`
	MetricsPath string `yaml:"metricspath"`
}

type raspberry struct {
	VcGenCmd string `yaml:"vcgencmd"`
}

var config *localConfig = nil

func Config() *localConfig {
	if config == nil {
		ParseConfig()
	}
	return config
}

func GetVcGenCmd() string {
	if _, err := os.Stat("/opt/vc/bin/vcgencmd"); !os.IsNotExist(err) {
		return "/opt/vc/bin/vcgencmd"
	} else if _, err := os.Stat("/usr/bin/vcgencmd"); !os.IsNotExist(err) {
		return "/usr/bin/vcgencmd"
	} else {
		panic("Could not find vcgencmd please install the raspberry pi toolchain")
	}
}

func ParseConfig() {
	flag.Parse()

	if configFileFlag := flag.Lookup("--config.file"); configFileFlag != nil {
		configFile := configFileFlag.Value.String()

		if configData, err := ioutil.ReadFile(configFile); err != nil {
			setDefaultConfig()
		} else if err = yaml.Unmarshal(configData, config); err != nil {
			setDefaultConfig()
		}
	} else {
		setDefaultConfig()
	}
}

func setDefaultConfig() {
	config = &localConfig{
		Raspberry: raspberry{
			VcGenCmd: GetVcGenCmd(),
		},
		Listen: listen{
			Address:     defaultListenAddress,
			MetricsPath: defaultMetricsPath,
		},
	}
}
