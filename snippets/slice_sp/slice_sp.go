package main

import "fmt"

func createSliceFromArray() {
	fmt.Println("=== createSliceFromArray ===")
	primes := [6]int{2, 3, 5, 7, 11, 13}

	var s []int = primes[1:4]
	fmt.Println(s)

	var s2 = s[0:2]
	fmt.Println(s2)

	s2[0] = 17 // slice 是索引，改变其中的一个其他的也都会发生改变
	fmt.Println(primes, s, s2)

	// s2[3] = 19 // panic: runtime error: index out of range
}

func createSliceFromLiterals() {
	fmt.Println("=== createSliceFromLiterals ===")
	q := []int{2, 3, 5, 7, 11, 13}
	fmt.Println(q)

	r := []bool{true, false, true, true, false, true}
	fmt.Println(r)

	s := []struct {
		i int
		b bool
	}{
		{2, true},
		{3, false},
		{5, true},
		{7, true},
		{11, false},
		{13, true},
	}
	fmt.Println(s)
}

func main() {
	createSliceFromArray()
	createSliceFromLiterals()
}
