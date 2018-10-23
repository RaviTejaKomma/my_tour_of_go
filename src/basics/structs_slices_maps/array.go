package main

import "fmt"

func main(){
	var a [2]string
	a[0] = "Mojo"
	a[1] = "Reads"
	fmt.Println(a[0],a[1])
	fmt.Println(a)

	primes := [6]int{2,3,5,7,11,13}
	fmt.Println(primes)

	var b [10]int
	for i:=0; i<10; i++{
		b[i] = i
	}
	fmt.Println(b)
}