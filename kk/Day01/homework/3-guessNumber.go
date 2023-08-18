package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	for {
		rand.Seed(time.Now().Unix())
		num := rand.Intn(100)
		fmt.Println(num)

		for times := 1; times <= 5; times++ {
			var guess int
			fmt.Println("猜数字游戏开始,五次机会,数字范围0-100")
			fmt.Printf("请输入数字:")
			fmt.Scan(&guess)
			if guess == num {
				fmt.Println("真牛，第", times, "就猜对了")
				break
			} else if guess > num {
				fmt.Println("太大了")
			} else {
				fmt.Println("太小了")
			}
			if times == 5 {
				fmt.Println("真xiu啊，5次都木架猜对")
			}
		}
		fmt.Println("是否继续？输入y yes继续")
		var input string
		fmt.Scan(&input)
		if input == "y" || input == "yes" {
			continue
		} else {
			break
		}

	}
}
