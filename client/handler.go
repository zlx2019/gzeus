/**
  @author: Zero
  @date: 2023/3/6 13:05:01
  @desc:

**/

package main

import (
	"fmt"
	"gzeus/network"
)

type MessageHandler func(packer *network.ClientPacket)

type InputHandler func(param *InputParam)

// 模拟的终端客户端请求处理函数
// 从这里转发请求到真正的游戏客户端

// Login 登录指令-处理函数
func (c *Client) Login(param *InputParam) {
	fmt.Println("接收到终端登录请求...")
	fmt.Println(param.Command)
	fmt.Println(param.Param)
}

func (c *Client) OnLoginRsp(packet *network.ClientPacket) {
}
func (c *Client) AddFriend(param *InputParam) {
}
func (c *Client) OnAddFriendRsp(packet *network.ClientPacket) {
}
func (c *Client) DelFriend(param *InputParam) {
}
func (c *Client) OnDelFriendRsp(packet *network.ClientPacket) {
}
func (c *Client) SendChatMsg(param *InputParam) {
}
func (c *Client) OnSendChatMsgRsp(packet *network.ClientPacket) {
}
