package main

import (
	"fmt"
	"math"
)

type MyFloat float64

func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

type Vertex struct {
	X, Y float64
}

func (v *Vertex) Abs() float64 { // 其实可以改写成 (v *Vertex)，不影响功能且减少拷贝
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

type Abser interface {
	Abs() float64
}

func main() {
	var a Abser

	// MyFloat 和 *Vertex 都实现了 Abs 方法，不用额外声明就实现了 Abser 接口
	f := MyFloat(-3.4)
	v := Vertex{3, 4}

	a = f
	fmt.Println(a.Abs())

	// a = v
	// cannot use v (type Vertex) as type Abser in assignment:
	// 		Vertex does not implement Abser (Abs method has pointer receiver)

	a = &v
	fmt.Println(a.Abs())
}
