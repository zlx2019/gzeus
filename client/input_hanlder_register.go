/**
  @author: Zero
  @date: 2023/3/6 13:49:49
  @desc:

**/

package main

import "gzeus/network/protocol/gen/messageId"

// InputHandlerRegister 根据枚举ID,注册对应的请求处理函数
func (c *Client) InputHandlerRegister() {
	c.inputHandlers[messageId.MessageId_LoginReq.String()] = c.Login
	c.inputHandlers[messageId.MessageId_AddFriendReq.String()] = c.AddFriend
	c.inputHandlers[messageId.MessageId_DelFriendReq.String()] = c.DelFriend
	c.inputHandlers[messageId.MessageId_SendChatMsgReq.String()] = c.SendChatMsg
}
