/*
- fmt.Print
- 不自动添加换行
- ,隔绝多个参数 自动用空格分隔
- 适合简单拼接输出

- fmt.Println
- 自动在结尾添加换行符（ \n ）
- 多个参数自动用空格分隔
- 适合结构化输出段落

- fmt.Printf
- 使用格式化占位符（如 %v 、 %T 、 %.2f ）
- 完全控制输出格式
- 需要手动添加换行符
- 适合需要精确控制的输出
*/

package main

import "fmt"

func main() {
	// Print 系列
	fmt.Print("My weight")        // 无换行，自动加空格
	fmt.Print(149.0*0.3783, "\n") // 连续输出会拼接：My weight149.0*0.3783...

	// Println 系列
	fmt.Println("My weight", 149.0*0.3783) // 自动加空格并换行
	fmt.Println("Age:", 41*365/687)        // 输出后会换行

	// Printf 格式化
	fmt.Printf("Weight: %.2f lbs\n", 149.0*0.3783) // 保留两位小数
	fmt.Printf("Age: %d years\n", 41*365/687)      // 整数格式化                              // 调用自
}
