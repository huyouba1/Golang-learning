package utils

import (
	"bufio"
	"crypto/md5"
	"fmt"
	"github.com/howeyc/gopass"
	"io"
	"math/rand"
	"os"
	"strings"
)

const (
	saltLength     = 6
	passwordLength = 6
)

const (
	normal = "普通用户"
	admin  = "管理员"
	root   = "超级管理员"
)

var (
	RoleList = []string{normal, admin, root}
	RoleMap  = map[string]string{
		"1": normal,
		"2": admin,
		"3": root,
	}
)

// 获取用户输入信息
func Input(prompt string) string {
	var text string
	fmt.Print(prompt)
	fmt.Scan(&text)
	return strings.TrimSpace(text)
}

// 判断文件是否存在
func FileIsExists(path string) bool {
	_, err := os.Stat(path)
	if err == nil {
		return true
	} else if os.IsNotExist(err) {
		return false
	} else {
		panic(err)
	}
}

// 读文件
func ReadFile(path string) string {
	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	txt := make([]byte, 0, 100)
	cxt := make([]byte, 10)
	for {
		n, err := file.Read(cxt)
		if err == io.EOF {
			break
		}
		txt = append(txt, cxt[:n]...)
	}
	return string(txt)
}

// 获取密码
func GetPassword(flag int) string {
	switch flag {
	case 0:
		fmt.Print("请设置密码: ")
	case 1:
		fmt.Print("请输入密码: ")
	case 3:
		fmt.Print("请输入新密码")
	}
	passwd, _ := gopass.GetPasswdMasked()
	return string(passwd)
}

// 生成salt
func generateSalt(limit int) string {
	chars := "qwertyuiopasdfghjklzxcvbnm1234567890QWERTYUIOPASDFGHJKLZXCVBNM"
	salt := make([]byte, limit)
	for i := 0; i < limit; i++ {
		salt[i] = chars[rand.Intn(len(chars))]
	}
	return string(salt)
}

// md5 Hash
func md5text(text string) string {
	return fmt.Sprintf("%x", md5.Sum([]byte(text)))
}

// 密码hash
func hashPassword(text, salt string) string {
	if salt == "" {
		salt = generateSalt(saltLength)
	}

	hash := md5text(fmt.Sprintf("%s:%s", salt, text))
	return fmt.Sprintf("%s:%s", salt, hash)
	//return md5text(fmt.Sprintf("%s:%s", salt, text))
	//return fmt.Sprintf("%s:%s", salt, text)
}

// 检查密码
func CheckPassword(text, hash string) bool {
	pos := strings.LastIndex(hash, ":")
	if pos < 0 {
		return false
	}
	return hashPassword(text, hash[:pos]) == hash
}

// 文件写操作函数
func WriteFile(path, txt string) {
	file, err := os.Create(path)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	writer := bufio.NewWriter(file)
	defer writer.Flush()
	writer.WriteString(txt)
}

// 设置密码
func SetPassword() string {
	for {
		password := GetPassword(1)
		if len(password) >= passwordLength {
			salt := generateSalt(saltLength)
			return hashPassword(password, salt)
		} else {
			fmt.Printf("密码长度不能少于%d位\n", passwordLength)
		}
	}
}

// 修改密码
//func ChangePassword(path string) {
//	switch Input("是否需要修改密码，请确认[y/yes]: ") {
//	case "y", "yes":
//		SetPassword(path)
//	default:
//		fmt.Println("取消修改密码！")
//	}
//}

// 设置用户权限
func SetRole() string {
	for {
		roleId := Input("请设置用户权限[1.普通用户 2.管理员 3.超级管理员]: ")
		if role, ok := RoleMap[roleId]; ok {
			return role
		} else {
			fmt.Println("输入错误!")
		}
	}
}
