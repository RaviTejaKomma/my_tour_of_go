package main

import "fmt"

/*
The zero value of a slice is nil.

A nil slice has a length and capacity of 0 and has no underlying array.
*/

func main(){
	var a []int
	fmt.Printf("len = %d, capacity = %d, %v\n",len(a),cap(a),a)
	if a == nil{
		fmt.Println("NIL")
	}
}