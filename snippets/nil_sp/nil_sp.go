package main

import "fmt"

type T struct {
	a string
	b int
}

func main() {
	fmt.Println(nil)

	var t T
	fmt.Println(t)
	fmt.Println(t.a) // 空字符串
	fmt.Println(t.b) // 0
}
