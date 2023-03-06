package player

import "gzeus/network/protocol/gen/messageId"

// HandlerRegister 将用户的基础方法绑定到 结构体的handlers属性中
func (player *Player) HandlerRegister() {
	player.handlers[messageId.MessageId_AddFriendReq] = player.AddFriend
	player.handlers[messageId.MessageId_DelFriendReq] = player.DelFriend
	player.handlers[messageId.MessageId_SendChatMsgReq] = player.ResolveChatMsg

}
