package main

import (
	"fmt"
	"time"
)

func main() {
	t1 := time.Date(2019, 5, 10, 17, 40, 0, 651387237, time.Local)
	t2 := time.Date(2019, 5, 10, 9, 40, 0, 651387237, time.UTC)
	fmt.Println(t1)
	fmt.Println(t2)

	fmt.Println(t1.Unix())
	fmt.Println(t2.Unix())

	fmt.Println(t1.Sub(t2)) // 0s
}
