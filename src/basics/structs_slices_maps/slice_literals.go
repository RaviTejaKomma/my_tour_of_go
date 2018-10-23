package main

import "fmt"

func main(){
	q := []int{2,3,5,7,11,13}
	fmt.Println(q)

	r := []bool{true, true, false, false, false, true}
	fmt.Println(r)
	s := []struct {
		 i int
		 b bool
	}{
		{2, true},
		{3, true},
		{5, false},
		{7, false},
		{11, false},
		{12, true},
	}
	fmt.Println(s)
}