package main

import "fmt"

func fibonacci(c, quit chan int) {
	x, y := 0, 1
	for {
		// 可以等待多个通信
		select {
		case c <- x: // 等待被读取
			x, y = y, x+y
		case <-quit: // 等待写入
			fmt.Println("quit")
			return
		}
	}
}

func main() {
	c := make(chan int)
	quit := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println(<-c) // 等待写入
		}
		quit <- 0
	}()
	fibonacci(c, quit)
}
