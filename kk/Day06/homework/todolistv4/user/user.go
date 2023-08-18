package user

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"todolistv4/utils"
)

var (
	PasswordFile  = "password.txt"
	PasswordLimit = 3
	TaskFile      = "task.txt"
)

type User struct {
	ID       int
	Name     string
	Password string
	Role     string
}

var Accounts = make([]*User, 0)

// init
func init() {
	loadUsers()
}

// 获取用户ID
func genUserId() int {
	var num int
	if len(Accounts) == 0 {
		return 1
	}
	for _, user := range Accounts {
		if num < user.ID {
			num = user.ID
		}
	}
	return num + 1
}

// new User对象
func NewUser() *User {
	return &User{
		ID:       genUserId(),
		Name:     "",
		Password: "",
	}
}

// 导入用户
func loadUsers() {
	if utils.FileIsExists(PasswordFile) {
		jsonContent, _ := ioutil.ReadFile(PasswordFile)
		err := json.Unmarshal(jsonContent, &Accounts)
		if err != nil {
			panic(err)
		}
	}
}

// 验证用户名密码
func VerifyPassword(username string, limit int) bool {
	for _, user := range Accounts {
		if username == user.Name {
			md5password := user.Password
			for count := 0; count < limit; count++ {
				input := utils.GetPassword(1)
				if utils.CheckPassword(input, md5password) {
					return true
				} else {
					fmt.Printf("密码验证错误，还剩%d次机会!\n", limit-count-1)
				}
			}
		}
	}
	return false
}

// 新增用户
func AddUser() {
	user := NewUser()
	for {
		name := utils.Input("请输入用户名：")
		if ValidateName(name) {
			user.Name = name
			break
		} else {
			fmt.Println("用户名格式有误！")
		}
	}
	password := utils.SetPassword()
	user.Password = password
	role := utils.SetRole()
	user.Role = role
	Accounts = append(Accounts, user)
	RecordAccounts()
	fmt.Println("账号添加成功")
}

// 验证用户名合法性（第一位必须是字母）
func ValidateName(txt string) bool {
	charsets := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	nums := "0123456789"
	for index, char := range txt {
		if index == 0 {
			if !strings.Contains(charsets, string(char)) {
				return false
			}
		} else {
			if !strings.Contains(charsets+nums, string(char)) {
				return false
			}
		}
	}
	return true
}

// 记录用户信息到json
func RecordAccounts() {
	ctx, err := json.Marshal(Accounts)
	if err != nil {
		panic(err)
	}
	file, _ := os.Create(PasswordFile)
	defer file.Close()

	var buffer bytes.Buffer
	json.Indent(&buffer, ctx, "", "\t")
	buffer.WriteTo(file)
}
