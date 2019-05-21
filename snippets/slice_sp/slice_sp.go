package main

import (
	"fmt"
	"github.com/binderclip/code-snippets-go/utils"
	"strings"
)

func createSliceFromArray() {
	utils.PrintFuncName()
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
	utils.PrintFuncName()
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
	utils.PrintFuncName()
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
	utils.PrintFuncName()
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
	utils.PrintFuncName()
	var s []int
	fmt.Println(s, len(s), cap(s))
	if s == nil {
		fmt.Println("nil!")
	}
}

func makingSlice() {
	utils.PrintFuncName()
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
	utils.PrintFuncName()
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
	utils.PrintFuncName()
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

func sliceAppend() {
	utils.PrintFuncName()
	var s []int
	printSlice(s)

	s1 := append(s, 1)
	printSlice(s1)

	s2 := append(s1, 2, 3, 4)
	printSlice(s2)

	s3 := s2[:2]
	printSlice(s3)

	s4 := append(s3, 5)
	s4[0] = 6
	printSlice(s4)
	printSlice(s2) // 发现 s4 append 的 5 是覆盖了 s2 中的 3

	s5 := append(s4, 7, 8, 9) // 超出了 cap
	s5[0] = 10
	printSlice(s5)
	printSlice(s2) // s2 已经和 s5 没有关系了
}

func sliceAppendSlice() {
	utils.PrintFuncName()
	s1 := []int{1, 2, 3}
	s2 := []int{100, 200, 300}
	s3 := append(s1, s2...)
	fmt.Println(s1)
	fmt.Println(s3)
}

func copySlice() {
	utils.PrintFuncName()
	s1 := []int{1, 2, 3}
	var s2 []int
	copy(s2, s1)
	fmt.Println(s2)

	s3 := make([]int, 2)
	copy(s3, s1)
	fmt.Println(s3)

	s4 := make([]int, 4)
	copy(s4, s1)
	fmt.Println(s4)
}

func rangeOfSlice() {
	utils.PrintFuncName()
	var pow = []int{1, 2, 4, 8, 16, 32, 64, 128}

	for i, v := range pow {
		fmt.Printf("2**%d = %d\n", i, v)
	}

	var pow2 = make([]int, 3)
	for i := range pow2 {
		pow2[i] = 1 << uint(i) // == 2**i
	}
	for _, v := range pow2 {
		fmt.Println(v)
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
	sliceAppend()
	sliceAppendSlice()
	copySlice()
	rangeOfSlice()
}
