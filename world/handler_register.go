/**
  @author: Zero
  @date: 2023/3/6 12:43:52
  @desc:

**/

package world

// HandlerRegister 注册函数
func (mm *MgrMgr) HandlerRegister() {
	mm.Handlers[1] = mm.UserLogin
}
