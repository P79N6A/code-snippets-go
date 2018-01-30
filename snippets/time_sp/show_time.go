package main

import "fmt"
import "time"

func main() {
	fmt.Println("Time is: ", time.Now())
	fmt.Println("Time is: ", time.Now().UnixNano())
}
