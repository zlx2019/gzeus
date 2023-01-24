package world

import (
	"gzeus/manager"
)

type MgrMgr struct {
	Pm manager.PlayerMgr
}

// NewMgrMgr 创建
func NewMgrMgr() *MgrMgr {
	return &MgrMgr{
		Pm: manager.PlayerMgr{},
	}
}

var MM *MgrMgr
