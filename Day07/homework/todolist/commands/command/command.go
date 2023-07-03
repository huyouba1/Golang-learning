package command

// 定义一个LoginCallBack，是一个函数类型，返回值是bool类型
// 如果返回true则登录成功，否则登录失败
type LoginCallBack func() bool

// 定义一个Callback，类型是函数类型
type CallBack func()

// 定义一个Command结构体
// 命令（选项）的名字和选择命令时调用的回调函数
type Command struct {
	Name     string
	CallBack CallBack
}

func NewCommand(name string, callback func()) *Command {
	return &Command{Name: name, CallBack: callback}
}
