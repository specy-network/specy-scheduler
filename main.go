package main

import (
	"fmt"
	"github.com/cosmos/relayer/v2/cmd"
	specyconfig "github.com/cosmos/relayer/v2/specy/config"
	"log"
	"path"
	"runtime"
)

func init() {
	if specyconfig.Config != nil {
		return
	}

	// 获取当前文件的路径
	_, filename, _, _ := runtime.Caller(0)
	root := path.Dir(filename)

	// 读取配置文件
	configPath := path.Join(root, "config.yaml")
	cfg, err := specyconfig.ReadConfigFile(configPath)
	if err != nil {
		log.Fatalf("Failed to read config file: %v", err)
		return
	}

	// 保存配置到全局变量
	specyconfig.Config = cfg

	fmt.Println("Chain Id:", specyconfig.Config.ChainId)
	fmt.Println("Chain Binary Location:", specyconfig.Config.ChainBinaryLocation)
	fmt.Println("Engine Node Address:", specyconfig.Config.EngineNodeAddress)
}

func main() {
	cmd.Execute()
}
