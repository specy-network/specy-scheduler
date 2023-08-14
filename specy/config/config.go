package config

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
	"path"
	"runtime"
)

var Config *SpecyConfig

type SpecyConfig struct {
	TargetChainId             string `yaml:"chain_id"`
	TargetChainBinaryLocation string `yaml:"chain_binary_location"`
	EngineNodeAddress         string `yaml:"engine_node_address"`
	HomeDir                   string `yaml:"home_dir"`
}

func ReadSpecyConfig() {
	// 获取当前文件的路径
	_, filename, _, _ := runtime.Caller(0)
	root := path.Dir(filename)

	// 读取配置文件
	configPath := path.Join(root, "../../config.yaml")
	cfg, err := readConfigFile(configPath)
	if err != nil {
		log.Fatalf("Failed to read config file: %v", err)
		return
	}

	// 保存配置到全局变量
	Config = cfg
}

func readConfigFile(filePath string) (*SpecyConfig, error) {
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
