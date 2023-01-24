package define

// HandlerParam 用户的方法名和方法参数
//
// HandlerKey 要执行的方法名
// Data		 执行方法携带的参数
type HandlerParam struct {
	HandlerKey string
	Data       interface{}
}
