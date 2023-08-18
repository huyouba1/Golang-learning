package main

import (
	"encoding/base64"
	"fmt"
)

func main() {

	// base64

	// 通常说的base64  0-9a-zA-Z+_  64个字符
	fmt.Println(base64.StdEncoding.EncodeToString([]byte("as看法空饭卡三森纳赛肯兰俩")))
	txt, _ := base64.StdEncoding.DecodeString("YXPnnIvms5Xnqbrppa3ljaHkuInmo67nurPotZvogq/lhbDkv6k=")
	fmt.Println(string(txt))

	// 在url中 +_ 是特殊字符，base64url(+(-)/(_))替换)
	fmt.Println(base64.URLEncoding.EncodeToString([]byte("as看法空饭卡三森纳赛肯兰俩")))
	txt, _ = base64.URLEncoding.DecodeString("YXPnnIvms5Xnqbrppa3ljaHkuInmo67nurPotZvogq_lhbDkv6k=")
	fmt.Println(string(txt))

	// 非对齐的方式 3 的整数倍 = 补齐
	// 标准
	fmt.Println(base64.RawStdEncoding.EncodeToString([]byte("as看法空饭卡三森纳赛肯兰俩")))
	fmt.Println(base64.RawURLEncoding.EncodeToString([]byte("as看法空饭卡三森纳赛肯兰俩")))
}
