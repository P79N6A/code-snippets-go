package main

import (
	"fmt"
	"github.com/binderclip/code-snippets-go/utils"
)

// Vertex 向量
type Vertex struct {
	X    int
	Y    int
	Name string
}

func vertex1() {
	utils.PrintFuncName()
	v1 := Vertex{1, 2, "a"}
	fmt.Println(v1)
	v1.X = 10
	fmt.Println(v1)
	fmt.Println(v1.X)

	v2 := Vertex{X: 11}
	fmt.Println(v2)

	v3 := Vertex{}
	fmt.Println(v3)
}

type T1 struct {
	f1 string
	f2 string
}

type T2 struct {
	T1
	f2     int64
	f3, f4 float64
}

func struct1() {
	utils.PrintFuncName()
	t := T2{T1{"foo", "bar"}, 1, 2, 3}
	fmt.Println(t.f1)
	fmt.Println(t.T1.f1)
	fmt.Println(t.f2, t.f3)
	fmt.Println(t)
}

func main() {
	vertex1()
	struct1()
}
