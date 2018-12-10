package main

import "fmt"

func main() {
	var a [2]string
	a[0] = "Hello"
	a[1] = "World"
	fmt.Println(a[0], a[1])
	fmt.Println(a)

	primes := [6]int{2, 3, 5, 7, 11, 13}
	fmt.Println(primes)

	// fmt.Println(primes[7]) // invalid array index 7 (out of bounds for 6-element array)

	// i := 7
	// fmt.Println(primes[i]) // panic: runtime error: index out of range
}
