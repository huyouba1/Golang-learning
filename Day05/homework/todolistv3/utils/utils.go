package utils

import (
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
func getPassword() string {
	fmt.Println("========== 欢迎使用huyouba1 todolist 系统 ===========")
	fmt.Print("请输入密码: ")
	password, _ := gopass.GetPasswdMasked()
	return string(password)
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
func checkPassword(text, hash string) bool {
	pos := strings.LastIndex(hash, ":")
	if pos < 0 {
		return false
	}
	return hashPassword(text, hash[:pos]) == hash
}

// 文件写操作函数
func WriteFile(path, txt string) {
	file, err := os.OpenFile(path, os.O_CREATE|os.O_APPEND|os.O_WRONLY|os.O_TRUNC, os.ModePerm)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	file.WriteString(txt)
}

// 设置密码,并将密码写到文件当中
func SetPassword(path string) {
	for {
		password := getPassword()
		if len(password) >= passwordLength {
			salt := generateSalt(saltLength)
			md5Password := hashPassword(password, salt)
			WriteFile(path, md5Password)
			fmt.Printf("密码设置成功，并成功存储到:%s\n", path)
			break
		} else {
			fmt.Printf("密码长度不能少于%d位\n", passwordLength)
		}
	}
}

// 修改密码
func ChangePassword(path string) {
	switch Input("是否需要修改密码，请确认[y/yes]: ") {
	case "y", "yes":
		SetPassword(path)
	default:
		fmt.Println("取消修改密码！")
	}
}

// 验证密码
func VerifyPasswd(path, password string, limit int) bool {
	hasher := ReadFile(path)
	for count := 0; count < limit; count++ {
		inputPasswd := getPassword()
		if checkPassword(inputPasswd, hasher) {
			return true
		} else {
			fmt.Printf("密码验证错误，还剩%d次机会!\n", limit-count-1)
		}
	}
	return false
}
