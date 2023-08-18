package main

import (
	"crypto/md5"
	"flag"
	"fmt"
	"math/rand"
	"strings"
	"time"
)

// 命令行工具
// -s  字符串计算hash => 自动随机的盐  => 结果输出(盐+hash)
// -c  检测密码  -s 密码原文  -hash  成功，失败
const saltLenth = 6

func init() {
	rand.Seed(time.Now().Unix())
}

func md5text(text string) string {
	return fmt.Sprintf("%x", md5.Sum([]byte(text)))
}

func checkPassword(text, hash string) bool {
	// text  + salt => texthash <=> hash
	// salt => hash $$ 之前的
	// 获取最后一个$
	pos := strings.LastIndex(hash, "$$")
	if pos < 0 {
		return false
	}

	return hashPassword(text, hash[:pos]) == hash
}

func generateSalt() string {
	chars := "qwertyuiopasdfghjklzxcvbnm1234567890QWERTYUIOPASDFGHJKLZXCVBNM"
	salt := make([]byte, saltLenth)
	for i := 0; i < saltLenth; i++ {
		salt[i] = chars[rand.Intn(len(chars))]
	}
	return string(salt)
}

func hashPassword(text, salt string) string {
	if salt == "" {
		salt = generateSalt()
	}
	hash := md5text(fmt.Sprintf("%s:%s", salt, text))
	// salt
	// 传递为空  随机生成salt(6位大小写英文字母数字组成)
	// salt$$hash
	return fmt.Sprintf("%s$$%s", salt, hash)
}

func main() {
	var (
		check bool
		text  string
		hash  string
		salt  string
		help  bool
		h     bool
	)

	flag.BoolVar(&h, "h", false, "help")
	flag.BoolVar(&help, "help", false, "help")
	flag.BoolVar(&check, "c", false, "check password")
	flag.StringVar(&text, "s", "", "text password")
	flag.StringVar(&salt, "salt", "", "salt password")
	flag.StringVar(&hash, "hash", "", "hash password")

	flag.Usage = func() {
		fmt.Println("Usage: [-c] check  [-s] text password  [-hash] hash")
		flag.PrintDefaults()
	}

	flag.Parse()

	if h || help {
		flag.Usage()
		return
	}

	if check {
		if ok := checkPassword(text, hash); ok {
			fmt.Println("密码正确")
		} else {
			fmt.Println("密码错误")
		}
	} else {
		fmt.Println("密码hash:", hashPassword(text, salt))
	}
}
