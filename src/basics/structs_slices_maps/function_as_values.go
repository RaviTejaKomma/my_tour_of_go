package main

import (
	"fmt"
	"math"
)

func compute(fn func(float64,float64) float64, x, y float64) float64{
	return fn(x,y)
}

func main(){
	hypotenuse := func(x, y float64) float64{
		return math.Sqrt(x*x + y*y)
	}
	fmt.Println(hypotenuse(5,12))
	fmt.Println(compute(hypotenuse,3,4))
	fmt.Println(compute(math.Pow,3,4))
}