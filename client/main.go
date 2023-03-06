/**
  @author: Zero
  @date: 2023/3/6 14:02:51
  @desc:

**/

package main

func main() {
	// 创建终端客户端
	client := NewClient()
	// 运行
	client.Run()
	select {}
}
