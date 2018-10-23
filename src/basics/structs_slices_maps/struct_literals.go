package main

import "fmt"

type Vertex struct{
	X int
	Y int
}

func main(){
	var v1 = Vertex{1,2}  // has type Vertex
	v2 := Vertex{X:1}  // Y:0 is implicit
	v3 := Vertex{}	   // X:0 and Y:0
	p := &Vertex{1,2}  // has type *Vertex
	fmt.Println(v1)
	fmt.Println(v2)
	fmt.Println(v3)
	fmt.Println(*p)
}