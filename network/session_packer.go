package network

// SessionPacket 服务端解包工具
// 用于解析来自于客户端的数据
type SessionPacket struct {
	Msg  *Message
	Sess *Session
}
