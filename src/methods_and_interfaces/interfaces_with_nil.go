package main

import "fmt"

type I interface{
	M()
}

type T struct{
	s string
}

func (t *T) M(){
	if t == nil{
		fmt.Println("<nil>")
		return
	}
	fmt.Println(t.s)
}

func describe(i I){
	fmt.Printf("Type: %T, Value: %v \n",i,i)
}

func main(){
	var i I
	var t *T
	i = t
	describe(i)
	i.M()

	i = &T{"MojoReads"}
	describe(i)
	i.M()
}