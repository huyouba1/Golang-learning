package main

import (
	"fmt"
	"os"
	"todolistv3/task"
	"todolistv3/utils"
)

var (
	passwordFile  = "password.txt"
	passwordLimit = 3
	taskFile      = "task.txt"
)

func main() {
	todolist := make([]task.Task, 0)
	if utils.FileIsExists(taskFile) {
		todolist = task.ReadTaskFromFile(taskFile)
	}

	if utils.FileIsExists(passwordFile) {
		password := utils.ReadFile(passwordFile)
		if utils.VerifyPasswd(passwordFile, password, passwordLimit) {
			fmt.Println("密码验证成功，请继续后续操作！")
		} else {
			fmt.Println("3次密码验证错误，程序退出！")
			os.Exit(1)
		}
	} else {
		utils.SetPassword(passwordFile)
	}

	for {
		text := utils.Input("请输入操作[1.添加任务 2.查询任务 3.修改任务 4.删除任务 5.修改密码 6.退出]:")
		if text == "6" {
			break
		}

		switch text {
		//add
		case "1":
			newTask := task.Add(todolist, passwordFile)
			todolist = append(todolist, newTask)
			task.RecordTask(taskFile, todolist...)
			// QueryTaskWithSort
		case "2":
			task.QueryTaskWithSort(todolist)
		// modify
		case "3":
			task.Modify(todolist, passwordFile)
			task.RecordTask(taskFile, todolist...)
		// remove
		case "4":
			if len(todolist) == 0 {
				fmt.Println("目前没有任何任务记录，请先添加任务，谢谢!")
				continue
			}
			task.Remove(todolist, passwordFile)
			task.RecordTask(taskFile, todolist...)
		case "5":
			utils.ChangePassword(passwordFile)
		case "6":
			break
		default:
			fmt.Println("输入的指令错误")
		}

	}
}
