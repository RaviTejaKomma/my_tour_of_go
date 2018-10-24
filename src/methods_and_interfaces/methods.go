package main

import(
	"fmt"
	"math"
)

/*
Go does not have classes. However, you can define methods on types.

A method is a function with a special receiver argument.

The receiver appears in its own argument list between the func keyword and the method name.

In this example, the Abs method has a receiver of type Vertex named v.
*/

type Vertex struct{
	x, y float64
}

func (v Vertex) Abs() float64{
	return math.Sqrt(v.x * v.x + v.y * v.y)
}

func main(){
	v := Vertex{3,4}
	fmt.Println(v.Abs())
}