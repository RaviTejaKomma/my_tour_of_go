package main

import "fmt"

/*
A slice is formed by specifying two indices, a low and high bound, separated by a colon:

a[low : high]

This selects a half-open range which includes the first element, but excludes the last one.
*/

func main(){
	primes := [6]int{2,3,5,7,11,13} 
	// or // 
	// var primes = [6]int{2,3,5,7,11,13}

	fmt.Println(primes)

	var s []int = primes[1:4]
	fmt.Println(s)
	s = primes[0:4]
	fmt.Println(s)
}