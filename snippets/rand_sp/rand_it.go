package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Println("n1", rand.Intn(100)) // 结果不变
	rand.Seed(time.Now().UnixNano())
	fmt.Println("n1", rand.Intn(100))
	fmt.Println("n2", rand.Intn(100))
}
