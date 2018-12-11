package main

import "fmt"

type Vertex struct {
	Lat, Long float64
}

func initMap() {
	fmt.Println("=== initMap ===")
	var m map[string]Vertex
	m = make(map[string]Vertex)
	fmt.Println(m)

	m["Bell Labs"] = Vertex{
		40.68433, -74.39967,
	}
	fmt.Println(m)
	fmt.Println(m["Bell Labs"])

	m2 := map[string]Vertex{
		"Bell Labs": Vertex{
			40.68433, -74.39967,
		},
		"Google": Vertex{
			37.42202, -122.08408,
		},
	}
	fmt.Println(m2)

	m3 := map[string]Vertex{
		"Bell Labs": {40.68433, -74.39967},
		"Google":    {37.42202, -122.08408},
	}
	fmt.Println(m3)
}

func mutatingMap() {
	fmt.Println("=== mutatingMap ===")
	m := make(map[string]int)

	m["Answer"] = 42
	fmt.Println("The Value:", m["Answer"])

	m["Answer"] = 37
	fmt.Println("The Value:", m["Answer"])

	delete(m, "Answer")
	fmt.Println("The Value:", m["Answer"])

	v, ok := m["Answer"]
	fmt.Println("The Value:", v, "Present?", ok)
}

func main() {
	initMap()
	mutatingMap()
}
