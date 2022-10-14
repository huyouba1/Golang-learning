package task

import "fmt"

// 包外可见
var Name = "Controller Task"
var Version string // 需要在运行启动之前进行初始化

// 属性
// 包外可见
// 包外是否可以修改 => 需要在包外不能修改
var version string

// 提供对外修改的函数（可读不可写）
func GetVersion() string {
	return version
}

func PrintVersion() {
	fmt.Println("v111111111111111")
}

func Call() {
	fmt.Println("controller Call")
}

// 包在使用的时候自动对这些信息进行初始化
func init() {
	Version = "Controller v1.0"
	fmt.Println("Controller init")
}
