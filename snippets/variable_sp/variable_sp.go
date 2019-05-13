package main

import "fmt"

var c, python, java bool

// variables with initializers
var c2, python2, java2 = true, false, "no!"

// short variable declarations，不能用在函数外面
// c3 := false // syntax error: non-declaration statement outside function body

func main() {
	var i int
	fmt.Println(i, c, python, java)

	var i2, j2 = 1, 2
	fmt.Println(i2, j2, c2, python2, java2)

	// short variable declarations，可以省略 var
	i3, j3 := 3, 4
	fmt.Println(i3, j3)
	// fmt.Println(c3)
}
