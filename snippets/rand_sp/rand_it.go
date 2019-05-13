package main

import (
	"fmt"
	"math/rand"
	"time"
)

const MaxRand = 100

func main() {
	fmt.Println("n1", rand.Intn(MaxRand)) // 结果不变
	rand.Seed(time.Now().UnixNano())
	fmt.Println("n1", rand.Intn(MaxRand))
	fmt.Println("n2", rand.Intn(MaxRand))
}
