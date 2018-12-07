package main

import "fmt"

func main() {
	// FILO 先进后出
	defer fmt.Println("!")
	defer fmt.Println("world") // defer 过的等最后 return 的时候才会真的执行

	fmt.Println("hello")
}
