// 烽火
package task

import (
	"bufio"
	"bytes"
	"encoding/csv"
	"encoding/gob"
	"encoding/json"
	"fmt"
	"github.com/olekukonko/tablewriter"
	"io/fs"
	"io/ioutil"
	"os"
	"path/filepath"
	"sort"
	"strconv"
	"strings"
	"time"
	"todolistv4/utils"
)

type Task struct {
	Id        int
	Name      string
	StartTime *time.Time
	EndTime   *time.Time
	Status    string
	User      string
}

var Todolist = make([]*Task, 0)

// init
func init() {
	loadTask()
}

// 导入task
func loadTask() {
	if utils.FileIsExists(JsonTaskFile) {
		jsonText, _ := ioutil.ReadFile(JsonTaskFile)
		err := json.Unmarshal(jsonText, &Todolist)
		if err != nil {
			panic(err)
		}
	}
}

// 结构体初始化函数
func NewTask() *Task {
	id := genId()
	now := time.Now()
	end := time.Now().Add(24 * time.Hour)
	return &Task{
		Id:        id,
		Name:      "",
		StartTime: &now,
		EndTime:   &end,
		Status:    statusNew, // 由于是new新任务，所以状态应该是"未执行"状态
		User:      "",
	}
}

// 获取最大ID
func genId() int {
	var rt int
	if len(Todolist) == 0 {
		return 1
	}

	for _, task := range Todolist {
		if rt < task.Id {
			rt = task.Id
		}
	}
	return rt + 1
}

// 验证状态值是否在可选列表内
func verifyStatus(inputStatus string) bool {
	for _, status := range StatusMap {
		if inputStatus == status {
			return true
		}
	}
	return false
}

// 验证任务名称是否唯一
func verifyName(inputName string) bool {
	for _, task := range Todolist {
		if inputName == task.Name {
			return false
		}
	}
	return true
}

// 表格渲染任务信息
func RenderTask(tasks ...*Task) {
	table := tablewriter.NewWriter(os.Stdout)
	table.SetAlignment(tablewriter.ALIGN_CENTER)
	table.SetHeader(Header)
	for i := 0; i < len(Header); i++ {
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

func time2string(t *time.Time) string {
	return t.Format(TimeLayout)
}

func ChangePassword() {
	switch utils.Input("是否需要修改密码[y/yes]: ") {
	case "y", "yes":
		utils.SetPassword()
	default:
		fmt.Println("取消修改密码！")
	}
}

func AddTask() {
	//var text string

	task := NewTask()

	fmt.Println("请输入任务信息：")
	for {
		tempName := utils.Input("任务名：")
		if verifyName(tempName) {
			task.Name = tempName
			break
		} else {
			fmt.Println("任务已存在！")
		}
	}
	task.User = utils.Input("负责人：")

	//utils.ChangePassword(path)
	//ChangePassword()
	Todolist = append(Todolist, task)
}

// 排序
func SortTask(tasks []*Task, key string) {
	if key == "name" {
		sort.Slice(tasks, func(i, j int) bool {
			return tasks[i].Name < tasks[j].Name
		})
	}
	if key == "startTime" {
		sort.Slice(tasks, func(i, j int) bool {
			return time2string(tasks[i].StartTime) < time2string(tasks[j].EndTime)
		})
	}
}

// 查询任务并排序
func QueryTaskWithSort() {
	var queryMap = map[string]string{"1": "name", "2": "startTime"}
	if len(Todolist) == 0 {
		fmt.Println("目前没有任何任务记录，请先添加任务，谢谢")
		return
	}

	filerTasks := make([]*Task, 0)
	q := utils.Input("请输入查询的任务名称: ")
	if q == "all" {
		filerTasks = Todolist
	} else {
		for _, task := range Todolist {
			if strings.Contains(task.Name, q) || strings.Contains(task.User, q) || strings.Contains(strconv.Itoa(task.Id), q) { // strings.Contains 判断是否包含字符串
				filerTasks = append(filerTasks, task)
			}
		}

	}
	if len(filerTasks) == 0 {
		fmt.Println("未找到关联任务!")
	} else {
		key := utils.Input("请输入排序方式[1.任务名称 2.任务开始时间]:")
		SortTask(filerTasks, queryMap[key])
		RenderTask(filerTasks...)
	}

}

// modify
func ModifyTask() bool {
	flag := false
	if len(Todolist) == 0 {
		fmt.Println("目前没有任何任务记录，请先添加任务，谢谢")
	}

	uid := utils.Input("请输入要修改的任务ID：")
	for idx, task := range Todolist {
		if id, _ := strconv.Atoi(uid); id == task.Id {
			RenderTask(task)
			switch utils.Input("是否确定修改？(y/Y/YES/yes)") {
			case "y", "Y", "yes", "YES":
				for {
					tempName := utils.Input("任务名称")
					if verifyName(tempName) {
						task.Name = tempName
						break
					} else {
						fmt.Println("任务名称已存在！")
					}
				}
				for {
					tempStatus := utils.Input("状态[1.未执行 2.开始执行 3.暂停 4.完成]:")
					if status, ok := StatusMap[tempStatus]; ok {
						task.Status = status
						now := time.Now()
						task.EndTime = &now
					} else {
						fmt.Println("输入的状态值不对！")
					}
				}
				Todolist[idx] = task
				flag = true
				//utils.ChangePassword(path)
				fmt.Println("任务修改完成！")
				RenderTask(task)
			default:
				fmt.Println("取消修改！")
				break
			}
		}
	}
	return flag
}

// delete
func DeleteTask() bool {
	flag := false
	queryID := utils.Input("请输入需要删除的任务ID：")
	for idx, task := range Todolist {
		if qid, _ := strconv.Atoi(queryID); qid == task.Id {
			RenderTask(task)
			switch utils.Input("是否进行删除(y/yes):") {
			case "y", "yes":
				copy(Todolist[idx:], Todolist[idx+1:])
				Todolist = Todolist[:len(Todolist)-1]
				//utils.ChangePassword(path)
				ChangePassword()
				fmt.Printf("任务ID:%s 已删除\n", queryID)
				flag = true
			default:
				fmt.Println("取消删除！")
				flag = false
			}
		}
	}
	return flag
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

func JsonTask() {
	ctx, err := json.Marshal(Todolist)
	if err != nil {
		panic(err)
	}
	file, _ := os.Create(JsonTaskFile)
	defer file.Close()

	var buffer bytes.Buffer
	json.Indent(&buffer, ctx, "", "\t")
	buffer.WriteTo(file)
}

func TxtTask() {
	content := make([]string, 0)
	for _, task := range Todolist {
		taskContent := []string{
			strconv.Itoa(task.Id),
			task.Name,
			time2string(task.StartTime),
			time2string(task.EndTime),
			task.Status,
			task.User,
		}
		content = append(content, strings.Join(taskContent, ","))
	}
	utils.WriteFile(TaskFile, strings.Join(content, "\n"))
}

// gob 持久化
func GobTask() {
	file, _ := os.Create(GobTaskFile)
	defer file.Close()
	encoder := gob.NewEncoder(file)
	for _, task := range Todolist {
		taskInfo := strings.Join([]string{strconv.Itoa(task.Id), task.Name, time2string(task.StartTime), time2string(task.EndTime), task.Status, task.User}, ", ")
		encoder.Encode(taskInfo)
	}
}

// CSV 持久化
func CsvTask() {
	taskFile := operateCsvFile()
	file, _ := os.Create(taskFile)
	defer file.Close()

	// 创建带缓冲IO的写对象
	writer := bufio.NewWriter(file)
	defer writer.Flush()

	csvWriter := csv.NewWriter(writer)
	csvWriter.Write(Header)
	for _, task := range Todolist {
		taskInfo := []string{
			strconv.Itoa(task.Id),
			task.Name,
			time2string(task.StartTime),
			time2string(task.EndTime),
			task.User,
		}
		csvWriter.Write(taskInfo)

	}

}

// 操作csv文件
func operateCsvFile() string {
	// 获取csv文件所在目录
	dir, _ := filepath.Split(CsvTaskFile)
	// 匹配csv文件
	pattern := dir + "tasks.csv*"
	matchFile := make([]string, 0)
	deleteFile := make([]string, 0)

	// 遍历目录下的文件
	filepath.Walk(dir, func(path string, info fs.FileInfo, err error) error {
		rt, _ := filepath.Match(pattern, path)
		if !info.IsDir() && rt {
			matchFile = append(matchFile, path)
		}
		return nil
	})
	// 对文件进行排序
	sort.Strings(matchFile)

	// 如果文件数量超过限制，则删除旧文件
	if len(matchFile) > RetainCsvNum-1 {
		deleteFile = matchFile[:len(matchFile)-2]
	}

	// 获取最新的文件
	latestfile := matchFile[len(matchFile)-1]
	fileinfo := strings.Split(latestfile, ".")
	lastfileSuffix, _ := strconv.Atoi(fileinfo[len(fileinfo)-1])
	suffix := strconv.Itoa(lastfileSuffix + 1)

	// 删除旧文件
	if len(deleteFile) > 0 {
		for _, path := range deleteFile {
			os.Remove(path)
		}
	}
	// 返回新文件名
	return strings.Join([]string{CsvTaskFile, suffix}, ".")
}

//// 保留最近N次的csv任务记录
//func operateCsvFile() string {
//	dir, _ := filepath.Split(CsvTaskFile)
//	pattern := dir + "tasks.csv*"
//	matchFile := make([]string, 0)
//	deleteFile := make([]string, 0)
//
//	filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
//		ret, _ := filepath.Match(pattern, path)
//		if !info.IsDir() && ret {
//			matchFile = append(matchFile, path)
//		}
//		return nil
//	})
//	sort.Strings(matchFile)
//
//	if len(matchFile) > RetainCsvNum-1 {
//		deleteFile = matchFile[:len(matchFile)-2]
//	}
//
//	latestfile := matchFile[len(matchFile)-1]
//	fileinfo := strings.Split(latestfile, ".")
//	lastfileSuffix, _ := strconv.Atoi(fileinfo[len(fileinfo)-1])
//	suffix := strconv.Itoa(lastfileSuffix + 1)
//
//	// delete old files
//	if len(deleteFile) > 0 {
//		for _, path := range deleteFile {
//			os.Remove(path)
//		}
//	}
//	return strings.Join([]string{CsvTaskFile, suffix}, ".")
//
//}
