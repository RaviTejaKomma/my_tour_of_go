package main

import (
	"fmt"
	"math"
)

/*
The expression T(v) converts the value v to the type T.

Some numeric conversions:

var i int = 42
var f float64 = float64(i)
var u uint = uint(f)
Or, put more simply:

i := 42
f := float64(i)
u := uint(f)
Unlike in C, in Go assignment between items of different type requires an explicit conversion. 
Try removing the float64 or uint conversions in the example and see what happens.
*/

func main(){
	var i, j int = 3, 4
	var k float64 = math.Sqrt(float64(i*i + j*j))
	var l uint = uint(k)
	fmt.Println(i,j,k,l) 
}