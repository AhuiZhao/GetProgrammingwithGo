package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	// 使用当前时间作为随机数种子，确保每次运行程序时生成的随机数不同
	localRand := rand.New(rand.NewSource(time.Now().UnixNano()))
	// 生成一个1-100的随机数
	randNum := localRand.Intn(100) + 1
	var guessNum int

	for {
		fmt.Print("请输入1-100的整数:")
		_, err := fmt.Scanln(&guessNum)
		// 检查输入是否有效
		if err != nil || guessNum < 1 || guessNum > 100 {
			fmt.Println("⚠️ 输入无效,请确保输入1-100的整数")
			continue
		}

		switch {
		case guessNum > randNum:
			fmt.Println("提示：猜大了")
		case guessNum < randNum:
			fmt.Println("提示：猜小了")
		default:
			fmt.Printf("🎉 恭喜！正确答案是 %d\n", randNum)
			return
		}
	}
}
