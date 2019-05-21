package main

import (
	"fmt"
	"github.com/binderclip/code-snippets-go/utils"
)

func array1() {
	utils.PrintFuncName()
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

	// fmt.Println(primes[7]) // invalid array index 7 (out of bounds for 6-element array)

	// i := 7
	// fmt.Println(primes[i]) // panic: runtime error: index out of range
}

func arrayModify1(arr [3]int) {
	arr[0] = 8
}

func arrayModify2(arr *[3]int) {
	arr[0] = 8
}

func array2() {
	utils.PrintFuncName()
	arr := [3]int{1, 2, 3}
	fmt.Println(arr)
	arrayModify1(arr)
	fmt.Println(arr)
	arrayModify2(&arr)
	fmt.Println(arr)
}

func main() {
	array1()
	array2()
}

// https://tour.golang.org/moretypes/6
// https://gobyexample.com/arrays
// slice 用的更多，array 定义的时候就 [] 中就确定了数字，slice 则没有
