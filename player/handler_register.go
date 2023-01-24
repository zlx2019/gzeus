package player

// HandlerRegister 将用户的基础方法绑定到 结构体的handlers属性中
func (player *Player) HandlerRegister() {
	player.handlers["add_friends"] = player.AddFriends
	player.handlers["del_friends"] = player.DelFriends
	player.handlers["resolve_chat_msg"] = player.ResolveChatMsg

}
