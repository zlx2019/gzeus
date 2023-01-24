package manager

import "gzeus/player"

// PlayerMgr 维护在线玩家
type PlayerMgr struct {
	// players 所有在线玩家容器 k:uid v:玩家
	players map[uint64]player.Player

	// 用于添加玩家的通道
	addPlayerChan chan player.Player
}

func NewPlayerMgr() PlayerMgr {
	return PlayerMgr{
		players:       make(map[uint64]player.Player),
		addPlayerChan: make(chan player.Player, 10),
	}
}

// Run 开始运行
func (pm *PlayerMgr) Run() {
	for {
		select {
		// 监听是否有新玩家上线
		case onlinePlayer := <-pm.addPlayerChan:
			pm.PlayerOnline(onlinePlayer)
		}
	}
}

// PlayerOnline 新玩家上线,添加到容器中
func (pm *PlayerMgr) PlayerOnline(p player.Player) {
	pm.players[p.UId] = p
}
