/**
  @author: Zero
  @date: 2023/3/6 13:56:03
  @desc:

**/

package main

import "gzeus/network/protocol/gen/messageId"

// MessageHandlerRegister 根据枚举ID,注册对应的响应处理函数
func (c *Client) MessageHandlerRegister() {
	c.messageHandlers[messageId.MessageId_LoginResp] = c.OnLoginRsp
	c.messageHandlers[messageId.MessageId_AddFriendResp] = c.OnAddFriendRsp
	c.messageHandlers[messageId.MessageId_DelFriendResp] = c.OnDelFriendRsp
	c.messageHandlers[messageId.MessageId_SendChatMsgResp] = c.OnSendChatMsgRsp
}
