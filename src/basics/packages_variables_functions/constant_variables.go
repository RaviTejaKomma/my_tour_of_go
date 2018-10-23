package main

import "fmt"

/*
Constants are declared like variables, but with the const keyword.

Constants can be character, string, boolean, or numeric values.

Constants cannot be declared using the := syntax.
*/

const Pi = 3.14

func main(){
	const i = 'R'
	const j = "mojo"
	const k = true
	fmt.Printf("Mr %c \n", i)
	fmt.Printf("Hello %v \n", j)
	fmt.Printf("Go rules? %v \n", k)
}