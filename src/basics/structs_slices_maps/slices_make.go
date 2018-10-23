package main

import "fmt"

func printSlice(s []int){
	fmt.Printf("len = %d, capacity = %d, %v\n",len(s),cap(s),s)
}

func main(){
	a := make([]int, 5)
	printSlice(a)

	b := make([]int, 0, 5)
	printSlice(b)

	c := b[:2]
	printSlice(c)

	d := b[2:5]
	printSlice(d)
}