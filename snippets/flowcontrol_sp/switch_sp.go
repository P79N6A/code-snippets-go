package main

import (
	"fmt"
	"runtime"
	"time"
)

func switchOne() {
	fmt.Println("=== switchOne ===")
	fmt.Print("Go runs on: ")
	// switch 的 case 都包含了 break
	// case 项不一定是常量和数字，可以放一个表达式
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("macOS.")
	case "linux":
		fmt.Println("Linux.")
	default:
		fmt.Printf("%s.", os)
	}
}

func switchTwo() {
	fmt.Println("=== switchTwo ===")
	today := time.Now().Weekday()
	// fmt.Println(time.Sunday)
	// fmt.Println(int(time.Sunday))
	// fmt.Println((today + 2) % 6)
	// switch 有执行顺序，前面的匹配了后面的就执行不到了
	switch time.Sunday {
	case today + 0:
		fmt.Println("Today.")
	case (today + 1) % 6:
		fmt.Println("Tomorrow.")
	case (today + 2) % 6:
		fmt.Println("In two days.")
	default:
		fmt.Println("Too far away.")
	}
}

func switchThree() {
	fmt.Println("=== switchThree ===")
	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("Good morning!")
	case t.Hour() < 17:
		fmt.Println("Good afternoon.")
	default:
		fmt.Println("Good evening.")
	}
}

func main() {
	switchOne()
	switchTwo()
	switchThree()
}
