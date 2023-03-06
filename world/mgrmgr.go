package world

import (
	"gzeus/manager"
	"gzeus/network"
)

// MgrMgr 服务器管理
type MgrMgr struct {
	Pm              *manager.PlayerMgr
	Server          *network.Server
	Handlers        map[uint64]func(message *network.SessionPacket)
	chSessionPacket chan *network.SessionPacket
}

// NewMgrMgr 创建
func NewMgrMgr() *MgrMgr {
	// 创建Mgr
	m := &MgrMgr{Pm: &manager.PlayerMgr{}}
	// 创建服务器
	m.Server = network.NewServer(":8023")
	m.Server.OnSessionPacket = m.OnSessionPacket
	return m
}

var MM *MgrMgr

// Run 运行服务器管理器
func (mm *MgrMgr) Run() {
	// 运行服务器
	go mm.Server.Run()
	// 运行用户管理器
	mm.Pm.Run()
}

func (mm *MgrMgr) OnSessionPacket(packet *network.SessionPacket) {
	if handler, ok := mm.Handlers[packet.Msg.ID]; ok {
		handler(packet)
	}
}
