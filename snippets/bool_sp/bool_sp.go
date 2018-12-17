package main

import "fmt"

func main() {
	var b bool
	b = true || false
	fmt.Println(b)
	b = true && false
	fmt.Println(b)
}
