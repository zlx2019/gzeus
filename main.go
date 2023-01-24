package main

import "gzeus/world"

// Main 启动程序入口
func main() {
	world.MM = world.NewMgrMgr()
	world.MM.Pm.Run()

	for {

	}
}
