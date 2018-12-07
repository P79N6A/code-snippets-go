package main

import "fmt"

// Vertex 向量
type Vertex struct {
	X int
	Y int
}

func main() {
	i, j := 42, 2701

	p := &i         // point to i
	fmt.Println(*p) // read i through the pointer
	*p = 21         // set i through the pointer
	fmt.Println(i)  // see the new value of i

	p = &j         // point to j
	*p = *p / 37   // divide j through the pointer
	fmt.Println(j) // see the new value of j

	v := Vertex{1, 2}
	p2 := &v
	p2.X = 420
	fmt.Println(v)
}
