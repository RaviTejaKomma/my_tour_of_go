package main

import "fmt"

/*
Struct is a collection of fields
Struct fields are accessed using a dot
*/

type Vertex struct{
	x int
	y int
}

func main(){
	fmt.Println(Vertex{1,2})
	v := Vertex{1,2}
	fmt.Println(v.x, v.y)
}