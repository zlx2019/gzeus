package network

import "net"

// ClientPacket 客户端解包工具
// 用来解析服务端数据包
type ClientPacket struct {
	Msg  *Message
	Conn net.Conn
}
