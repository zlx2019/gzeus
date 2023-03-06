package network

import "sync"

// SessionMgr 连接会话管理
type SessionMgr struct {
	// 连接的会话 key: 会话的UId val:会话
	Sessions map[int64]*Session
	// 计数器
	Counter int64
	// 锁
	Mutex sync.Mutex
	Pid   int64
}

var (
	SessionMgrInstance SessionMgr // 会话管理器单例实例
	onceInitSessionMgr sync.Once  // 保证只初始化会话管理单例一次
)

// 初始化
func init() {
	onceInitSessionMgr.Do(func() {
		// 实例化全局会话管理器
		SessionMgrInstance = SessionMgr{
			Sessions: make(map[int64]*Session),
			Counter:  0,
			Mutex:    sync.Mutex{},
		}
	})
}

// AddSession 将连接会话添加到会话管理器
func (mgr *SessionMgr) AddSession(s *Session) {
	// 加锁
	mgr.Mutex.Lock()
	// 保证最后解锁
	defer mgr.Mutex.Unlock()
	// 确保会话不存在在管理器才添加
	if val := mgr.Sessions[s.UId]; val != nil {
		if val.IsClose {
			mgr.Sessions[s.UId] = s
		} else {
			return
		}
	}
}

// DelSession 将会话从管理器中删除
func (mgr *SessionMgr) DelSession(UId int64) {
	delete(mgr.Sessions, UId)
}
