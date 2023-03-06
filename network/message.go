package network

// Message 客户端与服务端数据的传输包
type Message struct {
	// 数据包ID
	ID uint64
	// 数据
	Data []byte
}
