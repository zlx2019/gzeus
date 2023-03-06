package network

import (
	"net"
)

// Server 通信服务器
type Server struct {
	// 客户端连接监听器
	tcpListener net.Listener
	// 解包函数
	OnSessionPacket func(packet *SessionPacket)
}

// NewServer 根据地址,创建一个Server服务器
func NewServer(address string) *Server {
	// 创建一个tcp的地址指针
	resolveTCPAddr, err := net.ResolveTCPAddr("tcp6", address)
	if err != nil {
		panic(err)
	}
	// 建立tcp连接
	tcpListener, err := net.ListenTCP("tcp6", resolveTCPAddr)
	if err != nil {
		panic(err)
	}
	// 将建立好的tcp连接,设置到服务中
	s := &Server{tcpListener: tcpListener}
	return s
}

// Run 运行服务器
func (server *Server) Run() {
	for {
		// 阻塞,等待客户端连接
		clientConn, err := server.tcpListener.Accept()
		if err != nil {
			if _, ok := err.(net.Error); ok {
				continue
			}
		}
		// 开启一个协程,处理客户端连接
		go func() {
			// 将连接包装成一个会话
			session := NewSession(clientConn)
			SessionMgrInstance.AddSession(session)
			// 运行会话
			session.Run()
			SessionMgrInstance.DelSession(session.UId)
		}()
	}
}
