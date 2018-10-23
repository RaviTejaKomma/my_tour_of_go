package main

import "fmt"

/*
Pointers to structs
Struct fields can be accessed through a struct pointer.

To access the field X of a struct when we have the struct pointer p we could write (*p).X. 
However, that notation is cumbersome, so the language permits us instead to write just p.X, without the explicit dereference.
*/

type Vertex struct{
	X int
	Y int
}

func main(){
	v := Vertex{1,2}
	var p *Vertex
	p = &v
	p.X = 1e9
	fmt.Println(*p)
	fmt.Println(v)
}