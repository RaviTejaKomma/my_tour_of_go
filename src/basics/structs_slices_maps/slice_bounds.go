package main

import "fmt"

/*
Slice defaults
When slicing, you may omit the high or low bounds to use their defaults instead.

The default is zero for the low bound and the length of the slice for the high bound.

For the array

var a [10]int
these slice expressions are equivalent:

a[0:10]
a[:10]
a[0:]
a[:]
*/

func main(){
	a := [10]int{0,1,2,3,4,5,6,7,8,9}
	s := a[1:4]
	fmt.Println(s)
	s = a[:4]
	fmt.Println(s)
	s = a[1:]
	fmt.Println(s)
	s = a[:]
	fmt.Println(s)
}