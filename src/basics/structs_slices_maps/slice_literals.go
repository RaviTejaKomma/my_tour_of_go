package main

import "fmt"

/*
A slice literal is like an array literal without the length.

This is an array literal:

[3]bool{true, true, false}
And this creates the same array as above, then builds a slice that references it:

[]bool{true, true, false}
*/

func main(){
	q := []int{2,3,5,7,11,13}
	fmt.Println(q)

	r := []bool{true, true, false, false, false, true}
	fmt.Println(r)
	s := []struct {
		 i int
		 b bool
	}{
		{2, true},
		{3, true},
		{5, false},
		{7, false},
		{11, false},
		{12, true},
	}
	fmt.Println(s)
}