package config

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

var Config *SpecyConfig

type SpecyConfig struct {
	ChainId             string `yaml:"chain_id"`
	ChainBinaryLocation string `yaml:"chain_binary_location"`
	EngineNodeAddress   string `yaml:"engine_node_address"`
	HomeDir             string `yaml:"home_dir"`
}

func ReadConfigFile(filePath string) (*SpecyConfig, error) {
	// 读取配置文件内容
	content, err := ioutil.ReadFile(filePath)
	if err != nil {
		return nil, err
	}

	// 解析配置文件为结构体
	var config SpecyConfig
	if err := yaml.Unmarshal(content, &config); err != nil {
		return nil, err
	}

	return &config, nil
}
