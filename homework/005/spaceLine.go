package main

import (
	"fmt"
	"math/rand"
)

// 定义航空公司常量
const (
	SpaceAdventures = "Space Adventures"
	SpaceX          = "SpaceX"
	VirginGalactic  = "Virgin Galactic"
)

// 定义行程类型常量
const (
	RoundTrip = "Round-trip"
	OneWay    = "One-way"
)

func main() {
	// 无需手动 Seed，Go 1.20+ 会自动使用安全的默认随机源

	// 定义列名
	fmt.Printf("%-20s %-5s %-10s %s\n", "Spaceline", "Days", "Trip type", "Price")
	fmt.Println("=================================================")

	// 生成 10 条数据
	for i := 0; i < 10; i++ {
		// 随机选择航空公司
		spaceline := getRandomSpaceline()

		// 随机生成速度（16-30 km/s）
		speed := rand.Intn(15) + 16 // 16-30

		// 计算单程天数：距离 / 速度 / (60*60*24)
		distance := 62100000.0 // 公里
		days := distance / (float64(speed) * 60 * 60 * 24)

		// 随机选择行程类型
		tripType := getRandomTripType()

		// 随机生成价格（单程：3600万-5000万，往返价格翻倍）
		price := rand.Intn(1401) + 3600 // 3600-5000（百万美元）
		if tripType == RoundTrip {
			price *= 2
		}

		// 输出数据
		fmt.Printf("%-20s %-5.0f %-10s $ %d\n", spaceline, days, tripType, price)
	}
}

// 随机选择航空公司
func getRandomSpaceline() string {
	spacelines := []string{SpaceAdventures, SpaceX, VirginGalactic}
	return spacelines[rand.Intn(len(spacelines))]
}

// 随机选择行程类型
func getRandomTripType() string {
	tripTypes := []string{RoundTrip, OneWay}
	return tripTypes[rand.Intn(len(tripTypes))]
}
