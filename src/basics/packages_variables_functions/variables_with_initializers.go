package main

import "fmt"

/*
A var declaration can include initializers, one per variable.

If an initializer is present, the type can be omitted; the variable will take the type of the initializer.
*/

var i, j, k int = 1, 2, 3

func main(){
	var c, python = 1, "ravi"
	var java = true
	fmt.Println(i, j, k, c, python, java)
}