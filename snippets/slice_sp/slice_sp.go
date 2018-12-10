package main

import (
	"fmt"
	"strings"
)

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

func sliceDefault() {
	fmt.Println("=== sliceDefault ===")
	s := []int{2, 3, 5, 7, 11, 13}
	fmt.Println(s)
	fmt.Println(s[0:6])
	// fmt.Println(s[0:10]) // slice bounds out of range
	fmt.Println(s[:])
	fmt.Println(s[1:])
	fmt.Println(s[:3])
	// 不能像 Python 那样用负数
	// fmt.Println(s[:-2]) // invalid slice index -2 (index must be non-negative)
}

func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}

func sliceLenCap() {
	fmt.Println("=== sliceLenCap ===")
	s := []int{2, 3, 5, 7, 11, 13}
	printSlice(s)

	s = s[:0]
	printSlice(s)

	s = s[:4]
	printSlice(s)

	s = s[1:5] // 到这里最前面的 2 已经找不到了，cap 也变小了 1
	printSlice(s)

	s = s[0:1]
	printSlice(s)
}

func nilSlice() {
	fmt.Println("=== nilSlice ===")
	var s []int
	fmt.Println(s, len(s), cap(s))
	if s == nil {
		fmt.Println("nil!")
	}
}

func makingSlice() {
	fmt.Println("=== makingSlice ===")
	a := make([]int, 5)
	printSlice(a)
	b := make([]int, 1, 5)
	printSlice(b)
	c := b[:2]
	printSlice(c)
	d := b[2:cap(b)]
	printSlice(d)
}

func sliceOfSlice1() {
	fmt.Println("=== sliceOfSlice1 ===")
	// Create a tic-tac-toe board.
	board := [][]string{
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
	}
	// The players take turns.
	board[0][0] = "X"
	board[2][2] = "O"
	board[1][2] = "X"
	board[1][0] = "O"
	board[0][2] = "X"

	for i := 0; i < len(board); i++ {
		fmt.Printf("%s\n", strings.Join(board[i], " "))
	}
}

func sliceOfSlice2() {
	fmt.Println("=== sliceOfSlice2 ===")
	// 长度是可以不同的
	sos := [][]string{
		[]string{">"},
		[]string{">", ">"},
		[]string{">", ">", ">"},
		[]string{">", ">"},
		[]string{">"},
	}
	for i := 0; i < len(sos); i++ {
		fmt.Println(sos[i])
	}
}

func main() {
	createSliceFromArray()
	createSliceFromLiterals()
	sliceDefault()
	sliceLenCap()
	nilSlice()
	makingSlice()
	sliceOfSlice1()
	sliceOfSlice2()
}
