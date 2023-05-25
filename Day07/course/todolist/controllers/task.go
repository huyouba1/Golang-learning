package controllers

import "fmt"

//func AddTask() {
//	// 用户输入  拆分 view
//	// 数据验证  service
//	// 存储 service -> store/orm
//  // 数据结构显示  拆分  view
//	fmt.Println("添加任务")
//}

type TaskController struct {
}

func (c *TaskController) Add() {
	panic("test")
	fmt.Println("添加任务")

}

func (c *TaskController) Quary() {
	fmt.Println("查询成功")
}

func (c *TaskController) Modify() {
	fmt.Println("修改成功")
}

func (c *TaskController) ModifyStatus() {
	fmt.Println("修改任务状态成功")
}

func (c *TaskController) Delete() {
	fmt.Println("删除任务")
}
