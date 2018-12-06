package main

import "fmt"

// Pi the Pi
const Pi = 3.14

func main() {
	const World = "世界"
	fmt.Println("Hello", World)
	fmt.Println("Happy", Pi, "Day")

	const Truth = true
	fmt.Println("Go rules?", Truth)

	// Truth = false // cannot assign to Truth
	const Truth2 bool = false
	fmt.Println("Python rules?", Truth2)

	// const Wrold2 := "world" // syntax error: unexpected := at end of statement
}
