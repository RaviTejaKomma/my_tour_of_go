package main

import "fmt"

/*
Go's if statements are like its for loops; 
the expression need not be surrounded by parentheses ( ) but the braces { } are required.
*/

func even_or_odd(x int) bool{
	if x%2 == 0{
		return true
	}
	return false
}

func main(){
	var n int
	fmt.Scanln(&n)
	if even_or_odd(n){
		fmt.Println(n,"is even")
	}else{
		fmt.Println(n,"is odd")
	}
}