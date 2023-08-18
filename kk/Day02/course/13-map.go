package main

import (
	"fmt"
)

func main() {
	// 每个同学的成绩
	// key => ID value => 成绩
	var scores map[string]float64
	fmt.Printf("%T,%#v\n", scores, scores)

	// 初始化
	// 字面量
	scores = map[string]float64{} //空的map
	fmt.Printf("%T,%#v\n", scores, scores)

	scores = map[string]float64{"22": 80, "23": 90, "37": 70}
	fmt.Printf("%T,%#v\n", scores, scores)

	// make
	scores = make(map[string]float64) // ==  map[string]float64{}
	fmt.Printf("%T,%#v\n", scores, scores)
	scores = map[string]float64{"22": 80, "23": 90, "37": 70}

	// 操作
	fmt.Println(len(scores))

	// key => value
	// 查找
	fmt.Println(scores["22"])
	fmt.Println(scores["xx"]) // 因为定义的value为float，检索不到的时候则为空值 0
	fmt.Println(scores["tt"])

	// 判断 key 是否存在
	v, ok := scores["yy"]
	fmt.Println(ok, v)

	v, ok = scores["22"]
	fmt.Println(ok, v)

	// 改
	scores["22"] = 100
	fmt.Println(scores)

	// 增
	scores["yy"] = 120
	fmt.Println(scores)

	// 删除
	delete(scores, "yy")
	fmt.Println(scores)

	delete(scores, "aa")
	fmt.Println(scores)

	// 遍历映射
	for v := range scores {
		fmt.Println(v, scores[v])
	}

	fmt.Println("****************")

	for k, v := range scores {
		fmt.Println(k, v)
	}

}
