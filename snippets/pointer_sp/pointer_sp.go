package main

import (
	"fmt"
)

// Vertex 向量
type Vertex struct {
	X int
	Y int
}

func swapVertex1(v Vertex) {
	v.X, v.Y = v.Y, v.X
}

func swapVertex2(v *Vertex) {
	v.X, v.Y = v.Y, v.X
}

func pointer1() {
	fmt.Println("=== pointer1 ===")
	i, j := 42, 2701

	p := &i         // point to i
	fmt.Println(*p) // read i through the pointer
	*p = 21         // set i through the pointer
	fmt.Println(i)  // see the new value of i

	p = &j         // point to j
	*p = *p / 37   // divide j through the pointer
	fmt.Println(j) // see the new value of j

	v := Vertex{1, 2}
	fmt.Println(v)
	v1 := v
	v1.X = 210
	fmt.Println(v1)
	fmt.Println(v)
	p2 := &v
	p2.X = 420
	fmt.Println(v)

	swapVertex1(v)
	fmt.Println(v)
	swapVertex2(&v)
	fmt.Println(v)
}

func swap1(a int, b int) (int, int) {
	return b, a
}

func swap2(a *int, b *int) {
	*a, *b = *b, *a
}

func pointer2() {
	fmt.Println("=== pointer2 ===")
	a, b := 1, 2
	fmt.Println(swap1(a, b))
	fmt.Println(a, b)
	swap2(&a, &b)
	fmt.Println(a, b)
}

func main() {
	pointer1()
	pointer2()
}
