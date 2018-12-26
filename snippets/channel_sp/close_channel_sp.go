package main

import (
	"fmt"
)

func fibonacci(n int, c chan int) {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		c <- x
		x, y = y, x+y
	}
	close(c)
}

func main() {
	c := make(chan int, 10)
	go fibonacci(cap(c), c)
	for i := range c {
		fmt.Println(i)
	}

	// 其他的一些测试
	a := <-c
	fmt.Println(a)
	b, ok := <-c
	fmt.Println(b, ok)
	// c <- 10 // panic: send on closed channel
}
