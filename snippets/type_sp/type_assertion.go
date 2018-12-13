package main

import "fmt"

func main() {
	var i interface{} = "hello"

	s := i.(string)
	fmt.Println(s)

	// f := i.(float64) // panic: interface conversion: interface {} is string, not float64
	f, ok := i.(float64)
	fmt.Println(f, ok)

	s, ok = i.(string)
	fmt.Println(s, ok)
}
