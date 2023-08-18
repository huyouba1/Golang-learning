package main

import (
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/hex"
	"fmt"
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().Unix())
}

func readString(n int) (string, string) {
	rt1 := make([]byte, 0, n)
	rt2 := make([]byte, n, n)
	// 定义取值范围
	//chars := []byte{'a', 'b', 'c', 'd', 'e', 'f', 'h', 'i', 'j', 'k', 'l', 'm', 'n', 'u', 'v', 'w', 'x', 'y', 'z'}

	chars := "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

	// 循环n次，每次生成随机数(切片范围内)，获取对应的字符
	for i := 0; i < n; i++ {
		rt1 = append(rt1, chars[rand.Intn(len(chars))])
		rt2[i] = chars[rand.Intn(len(chars))]

	}
	return string(rt1), string(rt2)
}

func md5String(text string, salt string) string {
	//salt+ ":"+text
	bytes := []byte(salt)
	bytes = append(bytes, ':')
	bytes = append(bytes, []byte(text)...)

	return fmt.Sprintf("%x\n", md5.Sum(bytes))
}

func main() {
	// hash 算法 => 签名（不可逆）
	// MD5, sha1, sha256, sha512
	fmt.Printf("%x\n", md5.Sum([]byte("我是kk")))

	hasher := md5.New()
	hasher.Write([]byte("我是"))
	hasher.Write([]byte("kk"))
	fmt.Println(hex.EncodeToString(hasher.Sum(nil)))

	// 加盐 + md5
	salt1, salt2 := readString(6)
	fmt.Println(salt1, salt2)
	fmt.Println(md5String("我是kk", salt1))

	/*
		md5 hash => 不可逆
		string (n:1) => md5

		md5 => string

		彩虹表(暴力破解) MD5表 => string

		A => md5 密码iamkk => abcdef
		B => md5 密码iamkk => abcdef

		某一天A被脱裤了 iamkk => abcdef ==> iamkk

		加盐
		md5(iamkk + salt) =>  salt + defg

		A => md5 密码iamkk => abcdef +> iamkk +xxx => hash
		B => md5 密码iamkk saltxxx => xyz
	*/

	fmt.Printf("%x\n", sha1.Sum([]byte("我是kk")))
	fmt.Printf("%x\n", sha256.Sum256([]byte("我是kk")))
	fmt.Printf("%x\n", sha512.Sum512([]byte("我是kk")))

	sha256Hasher := sha256.New()
	sha256Hasher.Write([]byte("我是"))
	sha256Hasher.Write([]byte("kk"))
	fmt.Println(hex.EncodeToString(sha256Hasher.Sum(nil)))
}
