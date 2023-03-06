package player

import (
	"gzeus/network"
	"gzeus/network/protocol/gen/messageId"
)

// Player 用户结构体
// UId  用户的唯一标识
// FriendList 用户的好友列表
//
//
type Player struct {
	UId uint64
	// 好友列表
	FriendList []uint64
	// 通过该通道,执行用户已绑定到 handlers 的方法。
	HandlerParamChan chan *network.SessionPacket
	// 将用户的方法绑定到这里。根据Key进行执行.
	// key使用Proto中定义的枚举类型,通过该枚举执行不同的方法
	handlers map[messageId.MessageId]Handler
	// 用户会话
	session *network.Session
}

// NewPlayer 创建一个用户
func NewPlayer() *Player {
	p := &Player{
		UId:        0,
		FriendList: make([]uint64, 100),
		handlers:   make(map[messageId.MessageId]Handler),
	}
	// 将Player的功能函数注册到自身的handlers属性中
	p.HandlerRegister()
	return p
}

// Run 用户开始游戏
func (player *Player) Run() {
	for {
		// 监听用户
		select {
		// 监听到本用户要执行的方法信息
		case handlerParam := <-player.HandlerParamChan:
			if fn, ok := player.handlers[messageId.MessageId(handlerParam.Msg.ID)]; ok {
				// 获取到要执行的方法,然后进行执行
				fn(handlerParam)
			}
		}
	}
}
