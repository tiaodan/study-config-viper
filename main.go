package main

import (
	"fmt"
	"log"
	"study-config-viper/config"
)

func main() {
	fmt.Print("Hello, World!")
	// 获取配置实例（首次调用时触发初始化）
	cfg := config.GetConfig()
	cfg2 := config.GetConfig()
	cfg3 := config.GetConfig()
	log.Println("config: ", cfg2, cfg3)

	// 使用配置
	log.Println("network.local_ip: ", cfg.Network.LocalIp)
	log.Println("network.local_port: ", cfg.Network.LocalPort)
	log.Println("network.remote_ip: ", cfg.Network.RemoteIp)
	log.Println("network.remote_port: ", cfg.Network.RemotePort)
}
