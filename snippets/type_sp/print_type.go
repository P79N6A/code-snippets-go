package main

import "fmt"

type Vertex struct {
	X int
	Y int
}

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
	anArray := []int{1, 2, 3}
	fmt.Printf("%T\n", &anArray)
	fmt.Printf("%T\n", Vertex{X:2, Y:3,})
	fmt.Printf("%T\n", &Vertex{X:2, Y:3,})
}

// https://golangcode.com/print-variable-type/
