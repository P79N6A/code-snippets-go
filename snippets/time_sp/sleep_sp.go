package main

import (
	"fmt"
	"time"
)

func main() {
	for i := 0; i < 5; i++ {
		time.Sleep(300 * time.Millisecond)
		fmt.Println(i)
	}
}
