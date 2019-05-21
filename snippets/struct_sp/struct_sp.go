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

func main() {
	vertex1()
}
