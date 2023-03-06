package network

import (
	"encoding/binary"
	"fmt"
	"net"
	"time"
)

// Client 游戏客户端
type Client struct {
	Address string
	packer  IPacker
	// 客户端用来接收消息的通道
	ChMsg chan *Message
	// 根据一个MessageId,执行不同的handler
	OnMessage func(packet *ClientPacket)
}

// NewClient 客户端创建
func NewClient(address string) *Client {
	return &Client{
		Address: address,
		packer:  &NormalPacker{Order: binary.BigEndian},
		ChMsg:   make(chan *Message, 1),
	}
}

func (c *Client) Run() {
	// 连接服务端
	conn, err := net.Dial("tcp6", c.Address)
	if err != nil {
		fmt.Println(err)
		return
	}
	// 循环写数据到客户端
	go c.Write(conn)
	// 读取服务端响应数据
	go c.Read(conn)
}

// 向服务端写数据
func (c *Client) Write(conn net.Conn) {
	// 定义一个定时器,每秒触发一次,发送一条消息到服务端
	ticker := time.NewTicker(time.Second * 1)
	for {
		select {
		// 每秒执行一次
		case <-ticker.C:
			// 将自定义消息发送到 客户端通道
			c.ChMsg <- &Message{
				ID:   10001,
				Data: []byte("Hello"),
			}
		// 发送消息到服务端
		case msg := <-c.ChMsg:
			c.send(conn, msg)
		}
	}
}

// 写数据到服务端
func (c *Client) send(conn net.Conn, message *Message) {
	// 设置写数据超时时间
	err := conn.SetWriteDeadline(time.Now().Add(time.Second))
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("Client Send: [ID:%d  message:%s]\n", message.ID, string(message.Data))
	bytes, err := c.packer.Pack(message)
	if err != nil {
		fmt.Println(err)
		return
	}
	_, err = conn.Write(bytes)
	if err != nil {
		fmt.Println(err)
		return
	}
}

// 读取服务端响应数据
func (c *Client) Read(conn net.Conn) {
	for {
		message, err := c.packer.UnPack(conn)
		if _, ok := err.(net.Error); err != nil && ok {
			fmt.Println(err)
			continue
		}
		// 设置解包
		c.OnMessage(&ClientPacket{
			Msg:  message,
			Conn: conn,
		})
		fmt.Printf("Server Response: [ID:%d  message:%s]\n", message.ID, string(message.Data))
	}
}
