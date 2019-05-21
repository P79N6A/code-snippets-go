package main

import "fmt"

func array1() {
	fmt.Println("=== array1 ===")
	var b [3]string
	b[1] = "ho"
	fmt.Println(b)

	var a [2]string
	a[0] = "Hello"
	a[1] = "World"
	fmt.Println(a[0], a[1])
	fmt.Println(a)

	primes := [6]int{2, 3, 5, 7, 11, 13}
	fmt.Println(primes)
	fmt.Println(len(primes))

	nums := []int{1, 3, 5}
	fmt.Println(nums)
	fmt.Println(len(nums))

	// fmt.Println(primes[7]) // invalid array index 7 (out of bounds for 6-element array)

	// i := 7
	// fmt.Println(primes[i]) // panic: runtime error: index out of range
}

func arrayModify(arr []int) {
	arr[0] = 8
}

func array2() {
	fmt.Println("=== array2 ===")
	arr := []int{1, 2, 3}
	fmt.Println(arr)
	arrayModify(arr)
	fmt.Println(arr)
}

func main() {
	array1()
	array2()
}
