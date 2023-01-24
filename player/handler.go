package player

import (
	"fmt"
	"gzeus/chat"
	"gzeus/function"
)

type Handler func(any)

// AddFriends 新增一个或者多个好友
func (player *Player) AddFriends(data any) {
	// 断言data是否为 []uint64
	if uIds, ok := data.([]uint64); ok {
		for _, uId := range uIds {
			isFriend := function.CheckInNumberSlice(uId, player.FriendList)
			if !isFriend {
				// 只有非好友关系才能添加为好友
				player.FriendList = append(player.FriendList, uId)
			}
		}
	}
}

// DelFriends 删除一个或多个好友
func (player *Player) DelFriends(data any) {
	if uIds, ok := data.([]uint64); ok {
		for _, uId := range uIds {
			player.FriendList = function.DelElsInSlice(uId, player.FriendList)
		}
	}
}

// ResolveChatMsg 回复消息
func (player *Player) ResolveChatMsg(data any) {
	if msg, ok := data.(chat.Msg); ok {
		fmt.Println(msg)
	}
	//TODO 收到消息 然后转发到客户端
}
