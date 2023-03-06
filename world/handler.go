/**
  @author: Zero
  @date: 2023/3/4 16:11:43
  @desc:

**/

package world

import (
	"gzeus/network"
	"gzeus/player"
)

// UserLogin 用户登录
func (mm *MgrMgr) UserLogin(message *network.SessionPacket) {
	// 创建一个用户
	player := player.NewPlayer()
	player.UId = 111
	player.HandlerParamChan = message.Sess.WriteCh
	message.Sess.IsPlayerOnline = true
	// 运行用户
	player.Run()
}
