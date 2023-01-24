package network

import (
	"fmt"
	"net"
)

// Server 通信服务器
type Server struct {
	// 客户端连接监听器
	listener net.Listener
	// 服务器地址
	address string
	// 通信协议
	network string
}

// NewServer 根据配置创建Server
func NewServer(address, network string) *Server {
	return &Server{
		listener: nil,
		network:  network,
		address:  address,
	}
}

// Run 服务器运行
func (s *Server) Run() {
	// 创建tcp的addr
	tcpAddr, err := net.ResolveTCPAddr(s.network, s.address)
	if err != nil {
		fmt.Println(err)
		return
	}
	// 启动tcp监听器
	tcpListener, err := net.ListenTCP(s.network, tcpAddr)
	if err != nil {
		fmt.Println(err)
		return
	}
	s.listener = tcpListener

	for {
		// 接受来自于客户端的TCP连接
		conn, err := s.listener.Accept()
		if err != nil {
			continue
		}
		// 启用协程处理连接
		go func() {
			// 为每个连接创建一个Session,进行读写处理
			session := NewSession(conn)
			session.Run()
		}()
	}
}
