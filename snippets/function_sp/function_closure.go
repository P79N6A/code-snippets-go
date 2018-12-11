package main

import "fmt"

func adder() func(int) int {
	sum := 0 // sum 的值将会存在与返回的函数中
	return func(x int) int {
		sum += x
		return sum
	}
}

func adderClosure() {
	fmt.Println("=== adderClosure ===")
	pos, neg := adder(), adder()
	for i := 0; i < 10; i++ {
		fmt.Println(
			pos(i),
			neg(-2*i),
		)
	}
}

func fibonacci() func() int {
	i := 0
	j := 1
	return func() int {
		oldi := i
		i += j
		j = oldi
		return oldi
	}
}

func fibonacciClosure() {
	fmt.Println("=== fibonacciClosure ===")
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}

func main() {
	adderClosure()
	fibonacciClosure()
}
