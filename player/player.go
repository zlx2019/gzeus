package player

import (
	"gzeus/define"
)

// Player 用户结构体
// UId  用户的唯一标识
// FriendList 用户的好友列表
//
//
type Player struct {
	UId        uint64
	FriendList []uint64
	// 通过该通道,执行用户已绑定到 handlers 的方法。
	HandlerParamChan chan define.HandlerParam
	// 将用户的方法绑定到这里,方便执行。
	handlers map[string]Handler
}

// NewPlayer 创建一个用户
func NewPlayer() *Player {
	return &Player{
		UId:        0,
		FriendList: nil,
	}
}

// Run 用户开始游戏
func (player *Player) Run() {
	for {
		// 监听用户
		select {
		// 监听到本用户要执行的方法信息
		case handlerParam := <-player.HandlerParamChan:
			if fn, ok := player.handlers[handlerParam.HandlerKey]; ok {
				// 获取到要执行的方法,然后进行执行
				fn(handlerParam.Data)
			}
		}
	}
}
