package main

import "fmt"

// Vertex 向量
type Vertex struct {
	X    int
	Y    int
	Name string
}

func main() {
	v := Vertex{1, 2, "a"}
	fmt.Println(v)
	v.X = 10
	fmt.Println(v)
	fmt.Println(v.X)

	v2 := Vertex{X: 11}
	fmt.Println(v2)
}
