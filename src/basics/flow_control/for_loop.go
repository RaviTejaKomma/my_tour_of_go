package main

import "fmt"

/*
Go has only one looping construct, the for loop.

The basic for loop has three components separated by semicolons:

the init statement: executed before the first iteration
the condition expression: evaluated before every iteration
the post statement: executed at the end of every iteration
The init statement will often be a short variable declaration, 
and the variables declared there are visible only in the scope of the for statement.

The loop will stop iterating once the boolean condition evaluates to false.
*/

func type1(){
	sum := 0
	var n int
	fmt.Scanln(&n)
	for i := 0; i < n; i++{
		sum+=i
	}
	fmt.Println(sum)
}

func type2(){
	sum := 1
	var n int
	fmt.Scanln(&n)
	for ; sum < n; {
		sum+=sum
	}
	fmt.Println(sum)
}

func type3(){
	sum := 1
	var n int
	fmt.Scanln(&n)
	for sum < n{
		sum+=sum
	}
	fmt.Println(sum)
}

func main(){
	type1()
	type2()
	type3()
}