package main

import (
	"fmt"
	"math"
)

/*
In Go, a name is exported if it begins with a capital letter. For example, Pizza is an exported name, as is Pi, which is exported from the math package.

When importing a package, you can refer only to its exported names. 
Any "unexported" names are not accessible from outside the package.

Run the code. Notice the error message.

To fix the error, rename math.pi to math.Pi and try it again.
*/

func main(){
	fmt.Println(math.pi)
}