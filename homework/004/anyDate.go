package main

import (
	"fmt"
	"math/rand"
	"time"
)

// 闰年判断函数
func isLeapYear(year int) bool {
	return (year%4 == 0 && year%100 != 0) || (year%400 == 0)
}

// 获取月份天数
func daysInMonth(year, month int) int {
	switch month {
	case 2:
		if isLeapYear(year) {
			return 29
		}
		return 28
	case 4, 6, 9, 11:
		return 30
	default:
		return 31
	}
}

func main() {
	rand.Seed(time.Now().UnixNano())

	year := rand.Intn(131) + 1970 // 1970-2100年
	month := rand.Intn(12) + 1
	day := rand.Intn(daysInMonth(year, month)) + 1

	fmt.Printf("随机日期：%d-%02d-%02d\n", year, month, day)
}
