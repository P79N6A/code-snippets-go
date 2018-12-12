package main

import "fmt"

func main() {
	var i interface{}
	describe(i)

	i = 42
	describe(i)

	i = "hello"
	describe(i)

	describe(3.3)
}

func describe(i interface{}) {
	fmt.Printf("(%v, %T)\n", i, i)
}
