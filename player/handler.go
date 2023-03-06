package player

import (
	"fmt"
	"google.golang.org/protobuf/proto"
	"gzeus/function"
	"gzeus/network"
	pb_player "gzeus/network/protocol/gen/player"
)

// Handler 自定义处理函数,用来处理Player相关功能.
type Handler func(*network.SessionPacket)

// AddFriend 添加好友处理方法
// packer 客户端数据信息
func (player *Player) AddFriend(packer *network.SessionPacket) {
	// 将客户端发送的数据,解析到proto结构中.
	req := &pb_player.AddFriedRequest{}
	// 解析到req中.
	err := proto.Unmarshal(packer.Msg.Data, req)
	if err != nil {
		return
	}
	if !function.CheckInNumberSlice(req.UId, player.FriendList) {
		// 只有非好友关系才能添加为好友
		player.FriendList = append(player.FriendList, req.UId)
	}
}

// DelFriend 删除好友
func (player *Player) DelFriend(packer *network.SessionPacket) {
	req := &pb_player.DelFriedRequest{}
	err := proto.Unmarshal(packer.Msg.Data, req)
	if err != nil {
		return
	}
	player.FriendList = function.DelElsInSlice(req.UId, player.FriendList)
}

// ResolveChatMsg 回复消息
func (player *Player) ResolveChatMsg(packer *network.SessionPacket) {
	req := &pb_player.SendChatMsgRequest{}
	err := proto.Unmarshal(packer.Msg.Data, req)
	if err != nil {
		return
	}
	fmt.Println(req.Msg.GetContent())
	//TODO 收到消息 然后转发到客户端
}
