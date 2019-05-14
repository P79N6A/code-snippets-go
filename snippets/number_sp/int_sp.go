package main

import "fmt"

func main() {
	fmt.Println(0xFF) // 255
	fmt.Println(077) // 63
	//fmt.Println(0b1111, 0B1111, 0o17, 0O17) // will support in 1.13
}

// https://gfw.go101.org/article/basic-types-and-value-literals.html
