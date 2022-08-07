package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	for {
		// 生成随机整数
		rand.Seed(time.Now().Unix())
		num := rand.Intn(100)
		//fmt.Println(num)

		// 控制台输入猜测的数字
		var guess int
		fmt.Println("猜数字游戏开始,五次机会,数字范围0-100")
		for times := 1; times <= 5; times++ { // 循环输入的次数
			fmt.Print("请输入数字:")
			fmt.Scan(&guess)
			if guess == num {
				fmt.Println("真牛，第", times, "次就猜对了")
				break
			} else if guess > num {
				fmt.Println("输入的太大了")
			} else {
				fmt.Println("输入的太小了")
			}

			if times == 5 && guess != num {
				fmt.Println("5次机会用完了,游戏失败")
			}
		}

		var game string
		fmt.Println("是否重新开始游戏?(输入Y重新开始,任意键退出):")
		fmt.Scan(&game)
		if game == "Y" || game == "y" || game == "YES" || game == "yes" {
			continue
		} else {
			break
		}
	}

}
