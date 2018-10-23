package main

import "fmt"

type Vertex struct{
	lat, long float64
}
var m map[string]Vertex

func main(){
	m = make(map[string]Vertex)
	m["Bell Labs"] = Vertex{
		40.68433, -74.39967,
	}
	fmt.Println(m["Bell Labs"])

	// var m1 map[string]int
	// m1 = make(map[string]int)
	// or //
	m1 :=make(map[string]int)

	m1["ravi"] = 1
	m1["teja"] = 2
	m1["komma"] = 3

	fmt.Println(m1)
}