package main

import (
	"fmt"
)

/*
Interface values
Under the hood, interface values can be thought of as a tuple of a value and a concrete type:

(value, type)
An interface value holds a value of a specific underlying concrete type.

Calling a method on an interface value executes the method of the same name on its underlying type.
*/

type I interface{
	M()
}

type MyFloat float64

func (f MyFloat) M(){
	fmt.Println(f)
}

type T struct{
	s string
}

func (t *T) M(){
	fmt.Println(t.s)
}

func describe(i I){
	fmt.Printf("Type: %T, Value: %v\n",i,i)
}

func main(){
	var i I
	i = &T{"MojoReads"}
	describe(i)
	i.M()

	i = MyFloat(9)
	describe(i)
	i.M()
}