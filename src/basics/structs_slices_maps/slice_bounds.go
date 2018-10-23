package main

import "fmt"

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