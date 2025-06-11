package main

import (
	"fmt"
	"math/rand"
	"time"
)

// isLeapYear 判断是否是闰年
// 闰年规则：
// 1. 能被4整除但不能被100整除，或者
// 2. 能被400整除
func isLeapYear(year int) bool {
	return (year%4 == 0 && year%100 != 0) || (year%400 == 0)
}

// daysInMonth 返回指定年份和月份的天数
// 参数：year - 年份，month - 月份(1-12)
// 返回值：该月的天数
func daysInMonth(year, month int) int {
	switch month {
	case 2: // 二月特殊处理，需要考虑闰年
		if isLeapYear(year) {
			return 29 // 闰年二月有29天
		}
		return 28 // 平年二月有28天
	case 4, 6, 9, 11: // 4月、6月、9月、11月有30天
		return 30
	default: // 其他月份都有31天
		return 31
	}
}

func main() {
	// 初始化随机数生成器
	// 使用当前时间的纳秒数作为种子，确保每次运行结果不同
	localRand := rand.New(rand.NewSource(time.Now().UnixNano()))

	// 循环生成10个随机日期
	for i := 0; i < 10; i++ {
		// 生成1970-2100之间的随机年份
		year := localRand.Intn(131) + 1970
		// 生成1-12之间的随机月份
		month := localRand.Intn(12) + 1
		// 根据年份和月份生成随机天数
		// daysInMonth函数确保不会生成无效的日期（如2月30日）
		day := localRand.Intn(daysInMonth(year, month)) + 1

		// 格式化输出日期，月份和天数保持两位数显示
		fmt.Printf("%d. %d-%02d-%02d\n", i+1, year, month, day)
	}
}
