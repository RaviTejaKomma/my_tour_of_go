package main

import "fmt"

/*
A function can return any number of results
*/

func swap(x, y string) (string, string){
	return y, x
}

func main(){
	fmt.Println(swap("Mojo", "Reads"))
}