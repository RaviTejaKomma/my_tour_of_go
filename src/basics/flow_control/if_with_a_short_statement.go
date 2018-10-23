package main

import (
	"fmt"
	"math"
)

/*
Like for, the if statement can start with a short statement to execute before the condition.

Variables declared by the statement are only in scope until the end of the if.

(Try using v in the last return statement.)

Variables declared inside an if short statement are also available inside any of the else blocks.
*/

func pow(x, n, lim float64) float64 {
	if v := math.Pow(x, n); v < lim{
		return v
	}
	return lim
}

func main(){
	fmt.Println(
		pow(3,2,10),
		pow(3,3,20),
	)
}