package network

import (
	"encoding/binary"
	"fmt"
	"net"
	"time"
)

// Session 会话
// 用来维护客户端连接.使用会话来读取连接传来的数据或写回数据.
type Session struct {
	UId            int64               //会话ID
	Conn           net.Conn            // 客户端连接
	IsClose        bool                // 连接是否关闭
	packer         IPacker             // 封包-解包
	WriteCh        chan *SessionPacket // 将数据写回客户端事件通道
	IsPlayerOnline bool                // 玩家是否在线
	//MessageHandler func(packer *SessionPacket) //消息处理
}

// NewSession 根据一个客户端连接,创建一个Session
func NewSession(conn net.Conn) *Session {
	return &Session{Conn: conn, packer: NewNormalPacker(binary.BigEndian), WriteCh: make(chan *SessionPacket, 1)}
}

// Run 运行连接会话,循环读取数据、写回数据
func (s *Session) Run() {
	go s.Read()
	go s.Write()
}

// Read 读取客户端连接的数据
func (s *Session) Read() {
	for {
		// 解包读取客户端数据
		message, err := s.packer.UnPack(s.Conn)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Printf("Server Read Message : [ID:%d Data:%s]\n", message.ID, string(message.Data))
		// TODO 收到连接数据后,模拟写回数据
		s.WriteCh <- &SessionPacket{Msg: &Message{Data: []byte("Helllo")}, Sess: s}
	}
}

// Write 写回数据到客户端连接中
func (s *Session) Write() {
	for {
		select {
		// 数据写回客户端事件
		case message := <-s.WriteCh:
			s.send(message)
		}
	}
}

// 将数据响应到客户端连接
func (s *Session) send(packet *SessionPacket) {
	// 设置响应数据超时时间
	err := s.Conn.SetWriteDeadline(time.Now().Add(time.Second))
	if err != nil {
		fmt.Println(err)
	}
	bytes, err := s.packer.Pack(packet.Msg)
	if err != nil {
		return
	}
	_, err = s.Conn.Write(bytes)
	if err != nil {
		fmt.Println(err)
	}
}
