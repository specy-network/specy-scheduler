package main

import (
	"fmt"
	"github.com/cosmos/relayer/v2/specy"
	specyconfig "github.com/cosmos/relayer/v2/specy/config"
	"log"
	"time"
)

func init() {
	// 读取配置文件
	configPath := "config.yaml"
	cfg, err := specyconfig.ReadConfigFile(configPath)
	if err != nil {
		log.Fatalf("Failed to read config file: %v", err)
	}

	// 保存配置到全局变量
	specyconfig.Config = cfg

	fmt.Println("Chain Id:", specyconfig.Config.ChainId)
	fmt.Println("Chain Binary Location:", specyconfig.Config.ChainBinaryLocation)
	fmt.Println("Engine Node Address:", specyconfig.Config.EngineNodeAddress)
}

func main() {
	//cmd.Execute()

	// test scheduler
	now := time.Now()
	startTime := time.Date(now.Year(), now.Month(), now.Day(), 10, 0, 0, 0, now.Location())
	specy.StartScheduler(nil, "test", startTime, 10*time.Second)

	// 等待一段时间后停止 goroutine
	time.Sleep(time.Minute * 5)
}
