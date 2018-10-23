package main

import "fmt"

/*
A function can take zero or more arguments.

In this example, add takes two parameters of type int.

Notice that the type comes after the variable name.

When two or more consecutive named function parameters share a type, you can omit the type from all but the last.
*/

func add(x , y int) int {
	return x + y 
}

func main(){
	fmt.Println(add(2,4))
}