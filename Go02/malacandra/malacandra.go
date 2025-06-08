package main

import (
	"fmt"
	"math"
)

func malacandra() {
	fmt.Printf("速度最低需要%.0f公里/小时", math.Ceil(56000000.0/(28*24)))
}

// 添加主函数作为程序入口
func main() {
	malacandra() // 调用自定义函数
}
