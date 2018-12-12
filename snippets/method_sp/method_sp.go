package main

import (
	"fmt"
	"math"
)

type Vertex struct {
	X, Y float64
}

func (v Vertex) Abs() float64 { // 其实可以改写成 (v *Vertex)，不影响功能且减少拷贝
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func Abs2(v Vertex) float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

type MyFloat float64

func (f MyFloat) Abs3() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

func (v *Vertex) Scale(f float64) {
	// 用 pointer receiver 可以改值，不会拷贝数据
	v.X = v.X * f
	v.Y = v.Y * f
}

func (v Vertex) Scale2(f float64) {
	// 会拷贝数据再更改
	v.X = v.X * f
	v.Y = v.Y * f
}

func Scale3(v *Vertex, f float64) {
	// 用 pointer 可以改值
	v.X = v.X * f
	v.Y = v.Y * f
}

func Scale4(v Vertex, f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func main() {
	fmt.Println(Vertex{3, 4}.Abs())
	fmt.Println(Abs2(Vertex{5, 12}))

	f := MyFloat(-3.4)
	fmt.Println(f.Abs3())

	v := Vertex{3, 4}
	fmt.Println(v)
	fmt.Println("= Scale")

	v.Scale(0.5) // 并不会报错，其实相当于是 (&v).Scale(0.5) 的简写
	fmt.Println(v)
	(&v).Scale(2)
	fmt.Println(v)

	fmt.Println("= Scale2")
	v.Scale2(2)
	fmt.Println(v)

	fmt.Println("= Scale3")
	Scale3(&v, 2)
	fmt.Println(v)

	fmt.Println("= Scale4")
	Scale4(v, 2)
	fmt.Println(v)

	fmt.Println("= point receiver Scale")
	p := &Vertex{5, 8}
	p.Scale(2)
	fmt.Println(p)
	p.Scale2(2) // 并不会报错，去吃相当于是 (*p).Scale2(2) 的简写
	fmt.Println(p)
	p.Scale2(2)
	fmt.Println(p)
	(*p).Scale2(2)
	fmt.Println(p)

	fmt.Println("= point func Scale")
	p2 := &Vertex{5, 8}
	Scale3(p2, 2)
	fmt.Println(p2)

	// Scale4(p2, 2) // cannot use p2 (type *Vertex) as type Vertex in argument to Scale4
	Scale4(*p2, 2)
	fmt.Println(p2)
}
