/**
  @author: Zero
  @date: 2023/3/6 12:48:25
  @desc: 终端命令解析工具

**/

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// ConsoleClient 终端命令客户端
type ConsoleClient struct {
	// 用来接收操作命令的通道
	chInput chan *InputParam
}

// InputParam 指令结构
type InputParam struct {
	Command string   //操作指令
	Param   []string //指令参数
}

// NewConsoleClient 创建终端命令式客户端
func NewConsoleClient() *ConsoleClient {
	return &ConsoleClient{}
}

// Run 运行终端客户端
func (c *ConsoleClient) Run() {
	// 创建一个读取流
	reader := bufio.NewReader(os.Stdout)
	fmt.Println("Console Client 启动成功...")
	for {
		// 循环读取终端输入的命令
		//line, err := reader.ReadString('\n')
		bytes, _, err := reader.ReadLine()
		if err != nil {
			fmt.Println("读取终端命令错误!")
			continue
		}
		//解析命令
		lines := strings.Split(string(bytes), " ")
		if len(lines) == 0 {
			fmt.Println("命令格式有误!")
			continue
		}
		// 获取指令 和 指令参数
		command := lines[0]
		param := lines[1:]
		// 构建为命令对象
		input := &InputParam{
			Command: command,
			Param:   param,
		}
		// 发送命令
		c.chInput <- input
	}
}
