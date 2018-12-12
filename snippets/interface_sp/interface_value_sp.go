package main

import (
	"fmt"
	"math"
)

type I interface {
	M()
}

type T struct {
	S string
}

func (t *T) M() {
	if t == nil {
		fmt.Println(nil)
		return
	}
	fmt.Println(t.S)
}

type F float64

func (f F) M() {
	fmt.Println(f)
}

func main() {
	var i I // interface value

	i = &T{"Hello"}
	describe(i)
	i.M()

	// nil pointer
	var t *T // 未初始化的指针是 nil，但如果是 struct 的话是有值的
	i = t
	describe(i)
	i.M()

	i = F(math.Pi)
	describe(i)
	i.M()

	// nil interface
	var i2 I
	describe(i2) // (<nil>, <nil>)
	// i2.M() // panic: runtime error: invalid memory address or nil pointer dereference
}

func describe(i I) {
	fmt.Printf("(%v, %T)\n", i, i)
}
