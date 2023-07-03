package commands

import (
	"fmt"
	"github.com/olekukonko/tablewriter"
	"os"
	"strconv"
	"todolist/commands/command"
	"todolist/utils/ioutils"
)

// 定义一个manager结构体，结构体的属性cmds是一个map，map内的key是command对应的选择项，value则是具体的command
type manager struct {
	cmds          map[int]*command.Command
	loginCallback command.LoginCallBack
}

// 定义一个New函数，返回manager实例类型指针
func NewMgr() *manager {
	return &manager{
		// 给属性cmd初始化
		cmds: make(map[int]*command.Command),
	}
}

var mgr = NewMgr()

// 定义一个提示（说明）函数，打印出选项数字和选项名字
func (mgr *manager) prompt() {
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"编号", "功能说明"})
	for i := 1; i <= len(mgr.cmds); i++ {
		table.Append([]string{strconv.Itoa(i), mgr.cmds[i].Name}) // itoa转换编号，后面map编号对应的值
	}
	table.Render()
}

func (mgr *manager) getFunc(key int) (command.CallBack, error) {
	// 获取到对应的回调函数并返回
	if cmd, ok := mgr.cmds[key]; ok {
		return cmd.CallBack, nil
	} else {
		return nil, fmt.Errorf("不存在此选项！")
	}
}

// 定义一个注册方法，函数功能：将LoginCallBack添加到mgr中
func (mgr *manager) registerLoginCallBack(callback command.LoginCallBack) {
	mgr.loginCallback = callback
}

// 定义一个注册函数，此函数时暴露在外，给包外调用的
func RegisterLoginCallBack(callback command.LoginCallBack) {
	mgr.registerLoginCallBack(callback)
}

// 定义一个注册方法，函数功能：将 Command添加到manager.cmds中
func (mgr *manager) register(name string, callback command.CallBack) {
	mgr.cmds[len(mgr.cmds)+1] = command.NewCommand(name, callback)
}

// 定义一个注册函数，主义：此函数时暴露在外，给包外调用的。实际依然调用的是程序内部的register函数
func Register(name string, callback command.CallBack) {
	mgr.register(name, callback)
}

// 根据用户的输入获取(getFunc())回调函数并执行
func (mgr *manager) run() {
	if mgr.loginCallback != nil {
		if !mgr.loginCallback() {
			return
		}
	}
	for {
		// 调用提示（说明）函数
		mgr.prompt()
		// 打印出说明后，提示用户输入选项(选项数字),调用utils.iouitls.input函数
		key, err := strconv.Atoi(ioutils.Input("请输入功能编号: "))
		if err != nil {
			// 当发生错误时则输出，调用自定义的output输出
			ioutils.Error("输入错误！！！")
			continue
		}
		// 如果err为空则说明用户输出的正确，那么则根据用户的输出取出对应的回调函数，并执行
		if callback, err := mgr.getFunc(key); err != nil {
			/*
				err 不为空时，调用自定义的ioutils.Error
				err.Error()返回的是字符串
			*/
			ioutils.Error(err.Error())
		} else {
			// 执行回调函数
			callback()
		}

	}
}

func Run() {
	mgr.run()
}
