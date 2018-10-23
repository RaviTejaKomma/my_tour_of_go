package main

import "fmt"

/*
Inside a function, the := short assignment statement can be used in place of a var declaration with implicit type.

Outside a function, every statement begins with a keyword (var, func, and so on) and so the := construct is not available.
*/

func main(){
	var i, j int = 0, 1
	k := 2
	c, python := 1, "ravi"
	var java = true
	fmt.Println(i,j,k,c,python,java)
}