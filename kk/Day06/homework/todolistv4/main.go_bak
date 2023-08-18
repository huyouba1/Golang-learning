package main

import (
	"crypto/md5"
	"fmt"
	"github.com/howeyc/gopass"
	"github.com/olekukonko/tablewriter"
	"io"
	"os"
	"sort"
	"strconv"
	"strings"
	"time"
)

// 做一个命令行的任务管理器
// 用户管理

// 1. 函数，输入输出，复合数据结构，基本数据类型
// 2. 了解流程(对数据的操作流程，增删改查)

// 1. 任务的输入(添加任务)
// 2. 任务列表(任务查询)
// 3. 任务修改
// 4. 任务删除
// 5. 详情

// 任务
// ID,任务名称，开始时间，结束时间，状态，负责人
// ID,name,start_time,end_time,status,user
// []map[key][string]

var todos = []map[string]string{
	{"id": "1", "name": "陪孩子散步", "startTime": "18:00", "endTime": "19:00", "status": statusNew, "user": "huyouba1"},
	{"id": "2", "name": "备课", "startTime": "12:00", "endTime": "1:00", "status": statusNew, "user": "huyouba1"},
	{"id": "4", "name": "复习", "startTime": "11:00", "endTime": "9:00", "status": statusNew, "user": "huyouba1"},
}

const (
	id        = "id"
	name      = "name"
	startTime = "startTime"
	endTime   = "endTime"
	status    = "status"
	user      = "user"
	salt      = "02P8bA"
	password  = "b0230d6ac1b3b3bcff28ace36d15ad5d" // hello

)

const (
	statusNew     = "未执行"
	statusCompele = "完成"
)

// 定义函数 input 获取用户输入信息
func input(prompt string) string {
	var text string
	fmt.Print(prompt)
	fmt.Scan(&text)
	return strings.TrimSpace(text)
}

// 获取最大ID
func genId() int {
	var rt int
	for _, todo := range todos {
		todoId, _ := strconv.Atoi(todo["id"])
		if rt < todoId {
			rt = todoId
		}
	}
	return rt + 1
}

// 表格渲染任务信息
func renderTask(tasks []map[string]string) {
	header := []string{id, name, startTime, endTime, status, user}
	table := tablewriter.NewWriter(os.Stdout)
	table.SetAlignment(tablewriter.ALIGN_CENTER)
	table.SetHeader(header)
	for i := 0; i < len(header); i++ {
		table.SetColMinWidth(i, 20)
	}

	for _, task := range tasks {
		table.Append([]string{task[id], task[name], task[startTime], task[endTime], task[status], task[user]})
	}
	table.Render()
}

func printTask(task map[string]string) {
	fmt.Println(strings.Repeat("-", 20))
	fmt.Println("ID:", task[id])
	fmt.Println("任务名:", task[name])
	fmt.Println("开始时间:", task[startTime])
	fmt.Println("完成时间:", task[endTime])
	fmt.Println("任务状态:", task[status])
	fmt.Println(strings.Repeat("-", 20))

}

// 新任务初始化设置
func newTask() map[string]string {
	// id 生成(用todos中最大的ID+1)
	task := map[string]string{}
	task[id] = strconv.Itoa(genId())
	task[name] = ""
	task[startTime] = ""
	task[endTime] = ""
	task[status] = statusNew
	task[user] = ""
	return task
}

func add() {
	//var text string

	task := newTask()

	fmt.Println("请输入任务信息：")
	//fmt.Print("任务名：")
	//fmt.Scan(&text)
	task[name] = input("任务名：")

	//fmt.Print("开始时间：")
	//fmt.Scan(&text)
	task[startTime] = input("开始时间：")

	//fmt.Print("负责人：")
	//fmt.Scan(&text)
	task[user] = input("负责人：")

	todos = append(todos, task)
	fmt.Println("创建任务成功")
}

// 排序
func sortTask(tasks []map[string]string, key string) []map[string]string {
	if key == "name" || key == "startTime" {
		sort.Slice(tasks, func(i, j int) bool {
			return tasks[i][key] < tasks[j][key]
		})
	}
	return tasks
}

// 查询任务并排序
func queryTaskWithSort() {
	var queryMap = map[string]string{"1": "name", "2": "startTime"}
	q := input("请输入查询信息: ")
	content := make([]map[string]string, 0)
	for _, todo := range todos {
		if q == "all" || strings.Contains(todo[name], q) {
			content = append(content, todo)
		}
	}
	if len(content) == 0 {
		fmt.Println("未找到关联任务!")
	} else {
		key := input("请输入排序方式[1.任务名称 2.任务开始时间]:")
		newTask := sortTask(content, queryMap[key])
		renderTask(newTask)
	}

}

// 查询任务
func query() {
	// all 显示全部
	q := input("请输入查询信息名称：")
	//var text string
	//fmt.Print("请输入查询信息：")
	//fmt.Scan(&text)
	for _, todo := range todos {
		if q == "all" || strings.Contains(todo[name], q) {
			printTask(todo)
		}
	}
}

// modify
func modify() {

	for _, tname := range todos {
		printTask(tname)
	}
	uid := input("请输入要修改的任务ID：")
	var todo map[string]string
	for _, task := range todos {
		if task[id] == uid {
			todo = task
		}
	}
	//fmt.Println(uid)
	if todo != nil {
		fmt.Println("要修改的用户信息：")
		printTask(todo)
		inputInfo := input("是否确定修改？(y/Y/YES/yes)")
		if inputInfo == "y" || inputInfo == "Y" || inputInfo == "yes" || inputInfo == "YES" {
			todo[name] = input("新任务名:")
			todo[startTime] = input("开始时间:")
			todo[status] = input("状态:")
			if todo[status] == "完成" {
				todo[endTime] = time.Now().Format("2006-01-02 15:15:04:05")
			}
			fmt.Println("修改成功!")
		} else {
			fmt.Println("输入指令有误 (y/Y/YES/yes)")
		}
	} else {
		fmt.Println("输入的任务ID不存在")
	}

}

// delete
func remove() {
	for _, tname := range todos {
		printTask(tname)
	}
	uid := input("请输入要删除的任务ID：")
	var todo map[string]string
	for index, task := range todos {
		if task[id] == uid {
			todo = task
			if todo != nil {
				inputInfo := input("是否确定删除？(y/Y/YES/yes)")
				if inputInfo == "y" || inputInfo == "Y" || inputInfo == "yes" || inputInfo == "YES" {
					copy(todos[index:], todos[index+1:])
					todos = todos[:len(todos)-1]
					fmt.Println("删除成功")
					fmt.Println(todos)

				} else {
					fmt.Println("删除操作取消")
					return
				}
			} else {
				fmt.Println("输入的任务ID不存在")
				return
			}
		}
	}

}

// detail
func check() {
	var result string
	taskName := input("请输入任务名称，检测是否存在:")
	for _, todo := range todos {
		if taskName == todo[name] {
			//fmt.Printf("%#v\n", todo[name])
			//fmt.Printf("%#v\n", taskName)
			//fmt.Println("任务已存在")
			result = "任务已存在"
			break
		} else {
			//fmt.Printf("%#v\n", todo[name])
			//fmt.Printf("%#v\n", taskName)
			//fmt.Println("任务bu存在")
			result = "任务bu存在"
		}
	}
	fmt.Println(result)
}

// md5 加盐
func saltMd5(input string) string {
	hasher := md5.New()
	io.WriteString(hasher, input)
	io.WriteString(hasher, salt)

	cryptoPasswd := fmt.Sprintf("%x", hasher.Sum(nil))
	return cryptoPasswd
}

// 非明文显示终端输入信息，获取信息
func getPassword() string {
	fmt.Println("========== 欢迎使用huyouba1 todolist 系统 ===========")
	fmt.Print("请输入密码: ")
	password, _ := gopass.GetPasswdMasked()
	return string(password)
}

// 验证密码
func verifyPasswd() bool {
	limit := 3
	for count := 0; count < limit; count++ {
		inputPasswd := getPassword()
		if password == saltMd5(inputPasswd) {
			return true
		} else {
			fmt.Printf("密码验证错误，还剩%d次机会!\n", limit-count-1)
		}
	}
	return false
}

func main() {
	methods := map[string]func(){
		"1": add,
		"2": queryTaskWithSort,
		"3": modify,
		"4": remove,
		"5": check,
	}
	if !verifyPasswd() {
		fmt.Println("3次密码验证错误，程序退出")
		os.Exit(1)
	}
	for {
		//text := input("请输入操作(add/query/exit/.../)：")
		//fmt.Println("请输入操作(add/query/exit/.../)：")
		//var text string
		//fmt.Scan(&text)
		text := input("请输入操作[1.添加任务 2.查询任务 3.修改任务 4.删除任务 5.检查任务是否存在 6.退出]:")
		if text == "6" {
			break
		}
		//   //判断函数是否存在
		if method, ok := methods[text]; ok {
			method() // == add()，query()
		} else {
			fmt.Println("输入指令不正确 ")
		}

	}
}
