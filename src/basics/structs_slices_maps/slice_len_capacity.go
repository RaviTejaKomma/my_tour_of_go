package main

import "fmt"

/*
A slice has both a length and a capacity.

The length of a slice is the number of elements it contains.

The capacity of a slice is the number of elements in the underlying array, counting from the first element in the slice.

The length and capacity of a slice's can be obtained using the expressions len(s) and cap(s).

You can extend a slice's length by re-slicing it, provided it has sufficient capacity. 
Try changing one of the slice operations in the example program to extend it beyond its capacity and see what happens.
*/

func printSlice(s []int){
	fmt.Printf("len = %d, capacity = %d, %v\n",len(s),cap(s),s)
}
func main(){
	a := []int{2,3,5,7,11,13}
	printSlice(a)

	// Slice the slice to give it zero length.
	s := a[:0]
	printSlice(s)

	// Extend its length.
	s = a[:4]
	printSlice(s)

	// Drop its first two values.
	s = s[2:]
	printSlice(s)

	s = a[2:4]
	printSlice(s)
}