package main

import (
	"fmt"
	"time"
)

/*
Switch cases evaluate cases from top to bottom, stopping when a case succeeds.
*/

func main(){
	fmt.Println("When is Saturday? ")
	switch today:= time.Now().Weekday(); time.Saturday{
	case today+0:
		fmt.Println("Today.")
	case today+1:
		fmt.Println("Tomorrow.")
	case today+2:
		fmt.Println("Day after tomorrow.")
	default:
		fmt.Println("Too far away.")
	}
}