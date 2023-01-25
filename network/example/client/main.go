package main

import "gzeus/network"

func main() {
	// 模拟游戏客户端  连接服务端Port 8023
	client := network.NewClient(":8023")
	client.Run()
	select {}
}
