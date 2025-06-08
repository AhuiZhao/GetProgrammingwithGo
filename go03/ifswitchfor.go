package main

import (
	"fmt"
	"time"
)

func main() {
	var room = "100"
	var count = 10
	for count > 0 {
		fmt.Println(count)
		time.Sleep(1 * time.Second)
		count--

	}
	fmt.Println("Lift off!")
	if 7%2 == 0 {
		fmt.Println("7 is even")
	} else {
		fmt.Println("7 is odd")
	}
	switch {
	case room == "100":
		fmt.Println("Room 100")
		fallthrough //不用判断下一个条件 直接执行下一个case
	case room == "200":
		fmt.Println("Room 200")
	default:
		fmt.Println("Unknown room")
	}

}
