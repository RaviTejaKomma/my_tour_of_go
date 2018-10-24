package main

import (
	"fmt"
	"math"
)

/*
Methods and pointer indirection (2)
The equivalent thing happens in the reverse direction.

Functions that take a value argument must take a value of that specific type:

var v Vertex
fmt.Println(AbsFunc(v))  // OK
fmt.Println(AbsFunc(&v)) // Compile error!
while methods with value receivers take either a value or a pointer as the receiver when they are called:

var v Vertex
fmt.Println(v.Abs()) // OK
p := &v
fmt.Println(p.Abs()) // OK
In this case, the method call p.Abs() is interpreted as (*p).Abs().
*/

type Vertex struct{
	x, y float64
}

func (v Vertex) Abs() float64 {
	return math.Sqrt(v.x*v.x + v.y*v.y)
}

func AbsFunc(v Vertex) float64{
	return math.Sqrt(v.x*v.x + v.y*v.y)
}

func main(){
	v := Vertex{3,4}
	fmt.Println(v.Abs())
	fmt.Println(AbsFunc(v))

	p := &v
	fmt.Println(p.Abs())
	fmt.Println(AbsFunc(*p))
}