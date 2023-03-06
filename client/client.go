/**
  @author: Zero
  @date: 2023/3/6 13:01:56
  @desc:

**/

package main

import (
	"fmt"
	"gzeus/network"
	"gzeus/network/protocol/gen/messageId"
)

// Client 使用终端模拟的游戏客户端
type Client struct {
	// 游戏真正的客户端
	cli *network.Client
	// 终端命令对应的处理函数容器
	inputHandlers map[string]InputHandler
	// 终端命令对应的响应处理函数容器
	messageHandlers map[messageId.MessageId]MessageHandler
	// 游戏终端客户端
	console *ConsoleClient
	// 接收终端命令的用到
	chInput chan *InputParam
}

// NewClient 创建客户端
func NewClient() *Client {
	client := &Client{
		cli:             network.NewClient(":8023"),
		inputHandlers:   map[string]InputHandler{},
		messageHandlers: map[messageId.MessageId]MessageHandler{},
		console:         NewConsoleClient(),
	}
	client.cli.OnMessage = client.OnMessage
	client.cli.ChMsg = make(chan *network.Message, 1)
	client.chInput = make(chan *InputParam, 1)
	client.console.chInput = client.chInput
	// 执行函数绑定
	client.InputHandlerRegister()
	client.MessageHandlerRegister()
	return client
}

// Run 开始运行客户端
func (c *Client) Run() {
	go func() {
		for {
			select {
			// 等待终端输出命令
			case input := <-c.chInput:
				fmt.Printf("cmd:%s,param:%v  <<<\t \n", input.Command, input.Param)
				// 获取输出的指令,然后获取该指令对应的处理函数
				// 如果指令正确,并且有对应的函数则进行执行
				handler, ok := c.inputHandlers[input.Command]
				if ok {
					// 执行处理函数
					handler(input)
				}
			}
		}
	}()
	go c.console.Run()
	//go c.cli.Run()
	//go c.cli.Run()
}

func (c *Client) OnMessage(packet *network.ClientPacket) {
	if handler, ok := c.messageHandlers[messageId.MessageId(packet.Msg.ID)]; ok {
		handler(packet)
	}
}
