package main

import "fmt"

// 指定类型，用类型的默认值初始化
var c1, python1, java1 bool

// variables with initializers
var c2, python2, java2 = true, false, "no!"

// short variable declarations，不能用在函数外面
// c3 := false // syntax error: non-declaration statement outside function body

//var c4 bool, python4 bool = false, true // syntax error: unexpected comma after top level declaration
var c4 bool = false // 完整的写法会提示类型可省略

// 也可以组团声明
var (
	lang, bornYear, compiled     = "Go", 2007, true
	announceAt, releaseAt    int = 2009, 2012
	createdBy, website       string
)

// 依赖关系影响初始化顺序
var x, y = a+1, 5         // 8 5
var a, b, c = b+1, c+1, y // 7 6 5

func main() {
	var i1 int
	fmt.Println(i1, c1, python1, java1)

	var i2, j2 = 1, 2
	fmt.Println(i2, j2, c2, python2, java2)

	// short variable declarations，可以省略 var
	i3, j3 := 3, 4
	fmt.Println(i3, j3)
	// fmt.Println(c3)

	// 必须至少有一个新的变量
	//i3, j3 := 5, 6 // no new variables on left side of :=
	i3, j4 := 5, 6
	fmt.Println(i3, j4)
	i4, _ := 7, 8
	fmt.Println(i4, j4)

	fmt.Println(c4)

	// 局部变量必须要被使用
	//a := "aa" // a declared and not used

	fmt.Println(x, y, a, b, c)

	e, f := 10, 20
	e, f = f+1, e+1 // 21 11，和 Python 里的逻辑类似，大概是右边都算出值之后再一起给左边赋值
	fmt.Println(e, f)
}

// https://gfw.go101.org/article/constants-and-variables.html
