package network

import (
	"encoding/binary"
	"io"
	"net"
	"time"
)

// NormalPacker 传输数据 封包与解包工具
type NormalPacker struct {
	Order binary.ByteOrder
}

// NewNormalPacker 创建函数
func NewNormalPacker(order binary.ByteOrder) *NormalPacker {
	return &NormalPacker{order}
}

// Pack 封装数据包
// Message 要传输的数据
// 将服务端要传给客户端的数据进行封装成包
// 数据包结构: {数据包字节长度}|{数据ID}|{数据}
func (np *NormalPacker) Pack(message *Message) ([]byte, error) {
	// 定义包
	// 包的容量 = 数据包长度8字节(uint64) + 数据ID8字节(uint64) + 数据切片大小
	buf := make([]byte, 8+8+len(message.Data))

	// 设置包的0-8索引位置,为数据包的长度大小
	np.Order.PutUint64(buf[:8], uint64(len(buf)))

	// 设置包的8-16索引位置,为数据的ID
	np.Order.PutUint64(buf[8:16], message.Id)

	// 将要传输的数据拷贝到包的后面空余位置
	copy(buf[16:], message.Data)
	return buf, nil
}

// UnPack 解数据包
// 将客户端传来的数据进行解包
func (np *NormalPacker) UnPack(reader io.Reader) (*Message, error) {
	// 断言一下tcp连接
	// 并且设置读取数据超时时间
	err := reader.(*net.TCPConn).SetReadDeadline(time.Now().Add(2 * time.Second))
	if err != nil {
		return nil, err
	}
	// 先读取数据包的总长度和ID
	buf := make([]byte, 8+8)
	_, err = io.ReadFull(reader, buf)
	if err != nil {
		return nil, err
	}

	// 获取数据包的总长度(8字节)
	totalSize := np.Order.Uint64(buf[:8])
	// 数据ID(8字节)
	id := np.Order.Uint64(buf[8:16])
	// 计算得出 Data的长度
	dataSize := totalSize - 8 - 8

	// 有了Data的长度,就可以解析Data的数据
	// 这样解析,哪怕传入过多的数据,也不存在粘包的情况
	dataBuf := make([]byte, dataSize)
	_, err = io.ReadFull(reader, dataBuf)
	if err != nil {
		return nil, err
	}
	return &Message{id, dataBuf}, nil
}
