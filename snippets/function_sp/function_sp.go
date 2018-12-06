package main

import "fmt"

func add(x int, y int) int {
	return x + y
}

// x 和 y 的类型相同，省略了 x 的类型
func minus(x, y int) int {
	return x - y
}

// multiple results
func swap(x, y string) (string, string) {
	return y, x
}

func swapInt(x, y int) (int, int) {
	return y, x
}

// named return values
func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}

func main() {
	fmt.Println(add(42, 13))
	fmt.Println(minus(32, 40))

	a, b := swap("foo", "bar")
	fmt.Println(a, b)
	c, d := swapInt(1, 2)
	fmt.Println(c, d)

	fmt.Println(split(17))
}
