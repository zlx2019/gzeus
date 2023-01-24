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
	// 客户端连接
	conn net.Conn
	// 将数据写回客户端事件通道
	chWrite chan *Message
	// 封包\解包工具
	packer *NormalPacker
}

// NewSession 根据一个客户端连接,创建一个Session
func NewSession(conn net.Conn) *Session {
	return &Session{conn: conn, packer: NewNormalPacker(binary.BigEndian), chWrite: make(chan *Message, 1)}
}

// Run 运行连接会话,循环读取数据、写回数据
func (s *Session) Run() {
	go s.Read()
	go s.Write()
}

// Read 读取客户端连接的数据
func (s *Session) Read() {
	// 设置读取超时时间
	err := s.conn.SetReadDeadline(time.Now().Add(time.Second))
	if err != nil {
		fmt.Println(err)
	}
	for {
		// 解包读取客户端数据
		message, err := s.packer.UnPack(s.conn)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Printf("Server Read Message : [ID:%d Data:%s]\n", message.Id, string(message.Data))
		// TODO 收到连接数据后,模拟写回数据
		s.chWrite <- &Message{Id: 10001, Data: []byte("Hi Zero!")}
	}
}

// Write 写回数据到客户端连接中
func (s *Session) Write() {
	// 设置写入超时时间
	err := s.conn.SetWriteDeadline(time.Now().Add(time.Second))
	if err != nil {
		fmt.Println(err)
	}
	for {
		select {
		// 数据写回客户端事件
		case message := <-s.chWrite:
			s.send(message)
		}
	}
}

// 将数据响应到客户端连接
func (s *Session) send(message *Message) {
	bytes, err := s.packer.Pack(message)
	if err != nil {
		return
	}
	_, err = s.conn.Write(bytes)
	if err != nil {
		fmt.Println(err)
	}
}
