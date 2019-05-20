package main

import "fmt"
import "time"

func main() {
	fmt.Println("Time is: ", time.Now())
	fmt.Println("Time is: ", time.Now().UnixNano())
	fmt.Println("Time is: ", time.Now().Unix())
	fmt.Println("Time is: ", time.Unix(1558345862, 0))
}
