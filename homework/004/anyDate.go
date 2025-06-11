package main

import (
	"fmt"
	"math/rand"
	"time"
)

func isLeapYear(year int) bool {
	return (year%4 == 0 && year%100 != 0) || (year%400 == 0)
}

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
	// 使用Go 1.20+推荐的本地随机数生成器
	localRand := rand.New(rand.NewSource(time.Now().UnixNano()))

	year := localRand.Intn(131) + 1970 // 1970-2100年
	month := localRand.Intn(12) + 1
	day := localRand.Intn(daysInMonth(year, month)) + 1

	fmt.Printf("%d-%02d-%02d\n", year, month, day)
}
