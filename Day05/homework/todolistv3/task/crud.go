// 烽火
package task

import (
	"bufio"
	"fmt"
	"github.com/olekukonko/tablewriter"
	"os"
	"sort"
	"strconv"
	"strings"
	"time"
	"todolistv3/utils"
)

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
	statusBegin   = "开始执行"
	statusPause   = "暂停"
)

type Task struct {
	Id        int
	Name      string
	StartTime time.Time
	EndTime   time.Time
	Status    string
	User      string
}

var (
	statusChoice = []string{statusNew, statusCompele, statusBegin, statusPause}
	header       = []string{"ID", "Name", "StartTime", "EndTime", "Status", "User"}
)

// Task 结构体方法定义
func (task *Task) SetName(name string) {
	task.Name = name
}

func (task *Task) SetStartTime(StartTime time.Time) {
	task.StartTime = StartTime
}

func (task *Task) SetEndTime(EndTime time.Time) {
	task.EndTime = EndTime
}

func (task *Task) SetStatus(Status string) {
	task.Status = Status
}

func (task *Task) SetUser(User string) {
	task.User = User
}

// 结构体初始化函数
func NewTask(tasks []Task) *Task {
	id := genId(tasks)
	now := time.Now()
	end := time.Now().Add(24 * time.Hour)
	return &Task{
		Id:        id,
		Name:      "",
		StartTime: now,
		EndTime:   end,
		Status:    statusNew, // 由于是new新任务，所以状态应该是"未执行"状态
		User:      "",
	}
}

// 获取最大ID
func genId(tasks []Task) int {
	var rt int
	if len(tasks) == 0 {
		return 1
	}

	for _, task := range tasks {
		if rt < task.Id {
			rt = task.Id
		}
	}
	return rt + 1
}

// 验证状态值是否在可选列表内
func verifyStatus(inputStatus string) bool {
	for _, status := range statusChoice {
		if inputStatus == status {
			return true
		}
	}
	return false
}

// 验证任务名称是否唯一
func verifyName(tasks []Task, inputName string) bool {
	for _, task := range tasks {
		if inputName == task.Name {
			return false
		}
	}
	return true
}

// 表格渲染任务信息
func RenderTask(tasks ...Task) {
	table := tablewriter.NewWriter(os.Stdout)
	table.SetAlignment(tablewriter.ALIGN_CENTER)
	table.SetHeader(header)
	for i := 0; i < len(header); i++ {
		table.SetColMinWidth(i, 20)
	}

	for _, task := range tasks {
		table.Append([]string{
			strconv.Itoa(task.Id),
			task.Name,
			task.StartTime.Format("2006/01/02 15:04:05"),
			task.EndTime.Format("2006/01/02 15:04:05"),
			task.Status,
			task.User,
		})
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

func Add(tasks []Task, path string) Task {
	//var text string

	task := NewTask(tasks)

	fmt.Println("请输入任务信息：")
	for {
		tempName := utils.Input("任务名：")
		if verifyName(tasks, tempName) {
			task.SetName(tempName)
			break
		} else {
			fmt.Println("任务已存在！")
		}
	}
	task.SetUser(utils.Input("负责人："))

	//utils.ChangePassword(path)
	RenderTask(*task)
	return *task
}

// 排序
func sortTask(tasks []Task, key string) []Task {
	if key == "name" {
		sort.Slice(tasks, func(i, j int) bool {
			return tasks[i].Name < tasks[j].Name
		})
	}
	if key == "startTime" {
		sort.Slice(tasks, func(i, j int) bool {
			return tasks[i].StartTime.Before(tasks[j].StartTime)
		})
	}
	return tasks
}

//func sortTask(tasks []map[string]string, key string) []map[string]string {
//	if key == "name" || key == "startTime" {
//		sort.Slice(tasks, func(i, j int) bool {
//			return tasks[i][key] < tasks[j][key]
//		})
//	}
//	return tasks
//}

// 查询任务并排序
func QueryTaskWithSort(tasks []Task) {
	var queryMap = map[string]string{"1": "name", "2": "startTime"}
	if len(tasks) == 0 {
		fmt.Println("目前没有任何任务记录，请先添加任务，谢谢")
		return
	}

	q := utils.Input("请输入查询的任务名称: ")
	content := make([]Task, 0)

	for _, todo := range tasks {
		if q == "all" || strings.Contains(todo.Name, q) { // strings.Contains 判断是否包含字符串
			content = append(content, todo)
		}
	}
	if len(content) == 0 {
		fmt.Println("未找到关联任务!")
	} else {
		key := utils.Input("请输入排序方式[1.任务名称 2.任务开始时间]:")
		newTask := sortTask(content, queryMap[key])
		RenderTask(newTask...)
	}

}

// modify
func Modify(tasks []Task, path string) {

	if len(tasks) == 0 {
		fmt.Println("目前没有任何任务记录，请先添加任务，谢谢")
		return
	}

	uid := utils.Input("请输入要修改的任务ID：")
	for idx, task := range tasks {
		if id, _ := strconv.Atoi(uid); id == task.Id {
			RenderTask(task)
			switch utils.Input("是否确定修改？(y/Y/YES/yes)") {
			case "y", "Y", "yes", "YES":
				for {
					tempName := utils.Input("任务名称")
					if verifyName(tasks, tempName) {
						task.SetName(tempName)
						break
					} else {
						fmt.Println("任务名称已存在！")
					}
				}
				for {
					tempStatus := utils.Input("状态:")
					if verifyStatus(tempStatus) {
						if tempStatus == statusCompele {
							task.SetEndTime(time.Now())
						}
						task.SetStatus(statusCompele)
						break
					} else {
						fmt.Println("输入的状态值不对！可选范围：", strings.Join(statusChoice, ","))
					}
				}
				tasks[idx] = task
				//utils.ChangePassword(path)
				fmt.Println("任务修改完成！")
				RenderTask(task)
			default:
				fmt.Println("取消修改！")
				break
			}
		}
	}
}

//
//var todo map[string]string
//for _, task := range todos {
//	if task[id] == uid {
//		todo = task
//	}
//}
////fmt.Println(uid)
//if todo != nil {
//	fmt.Println("要修改的用户信息：")
//	printTask(todo)
//	inputInfo := input("是否确定修改？(y/Y/YES/yes)")
//	if inputInfo == "y" || inputInfo == "Y" || inputInfo == "yes" || inputInfo == "YES" {
//		todo[name] = input("新任务名:")
//		todo[startTime] = input("开始时间:")
//		todo[status] = input("状态:")
//		if todo[status] == "完成" {
//			todo[endTime] = time.Now().Format("2006-01-02 15:15:04:05")
//		}
//		fmt.Println("修改成功!")
//	} else {
//		fmt.Println("输入指令有误 (y/Y/YES/yes)")
//	}
//} else {
//	fmt.Println("输入的任务ID不存在")
//}

//}

// delete
func Remove(tasks []Task, path string) []Task {
	newTasks := make([]Task, 0)
	queryID := utils.Input("请输入需要删除的任务ID：")
	for idx, task := range tasks {
		if qid, _ := strconv.Atoi(queryID); qid == task.Id {
			RenderTask(task)
			switch utils.Input("是否进行删除(y/yes):") {
			case "y", "yes":
				copy(tasks[idx:], tasks[idx+1:])
				newTasks = tasks[:len(tasks)-1]
				//utils.ChangePassword(path)
				fmt.Printf("任务ID:%s 已删除\n", queryID)
			}
		}
	}
	return newTasks
}

// 持久化任务信息
func RecordTask(path string, tasks ...Task) {
	content := make([]string, 0)
	for _, task := range tasks {
		taskContent := []string{
			strconv.Itoa(task.Id),
			task.Name,
			task.StartTime.Format("2006/01/02 15:04:05"),
			task.EndTime.Format("2006/01/02 15:04:05"),
			task.Status,
			task.User,
		}
		content = append(content, strings.Join(taskContent, ","))
	}
	utils.WriteFile(path, strings.Join(content, "\n"))
}

// 从文件读取task信息
func ReadTaskFromFile(path string) []Task {
	var taskRecords []Task
	var task Task

	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		taskSlice := strings.Split(scanner.Text(), ",")
		id, _ := strconv.Atoi(taskSlice[0])
		startTime, _ := time.Parse("2006/01/02 15:04:05", taskSlice[2])
		endTime, _ := time.Parse("2006/01/02 15:04:05", taskSlice[3])
		task = Task{id, taskSlice[1], startTime, endTime, taskSlice[4], taskSlice[5]}
		taskRecords = append(taskRecords, task)
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}
	return taskRecords
}
