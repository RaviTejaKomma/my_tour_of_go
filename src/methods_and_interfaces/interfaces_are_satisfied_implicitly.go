package main

import "fmt"

type T struct{
	s string
}

type I interface{
	M()
}

func (t T) M(){
	fmt.Println(t.s)
}

func main(){
	// t := T{"Golang"}
	// var i I
	// i = t
	// i.M()

	// or //
	var i I = T{"Golang"}
	i.M()
}