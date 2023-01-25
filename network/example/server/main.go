package main

import "gzeus/network"

func main() {
	// 模拟服务端 端口为8023,协议为tcp6
	server := network.NewServer(":8023", "tcp6")
	server.Run()
	select {}
}
