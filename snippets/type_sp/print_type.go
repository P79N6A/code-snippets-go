package main

import "fmt"

func main() {
	anInt := 1234
	fmt.Printf("%T\n", anInt)
	fmt.Printf("%T\n", 123)
	fmt.Printf("%T\n", 123.3)
	fmt.Printf("%T\n", "abc")
	fmt.Printf("%T\n", 'a')
	fmt.Printf("%T\n", []int{1, 2, 3})
	fmt.Printf("%T\n", []int{1, 2, 3}[:1])
	fmt.Printf("%T\n", make([]int, 2))
}

// https://golangcode.com/print-variable-type/