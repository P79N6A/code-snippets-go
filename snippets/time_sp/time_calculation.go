package main

import (
	"fmt"
	"time"
)

func main() {
	t1 := time.Now()
	t2 := time.Date(2019, 5, 10, 17, 40, 0, 651387237, time.Local)

	fmt.Println(t1)
	fmt.Println(t2)
	fmt.Println(int(t2.Sub(t1) / time.Hour / 24))
}
