package main

import (
	"fmt"
	"strings"
)

func main(){
	m := make(map[string]int)
	m["Number"] = 42
	fmt.Println("The value is :",m["Number"])

	m["Number"] = 48
	fmt.Println("The value is :", m["Number"])

	delete(m,"Number")
	fmt.Println("The value is :", m["Number"])

	v, ok := m["Number"]
	fmt.Println("The value:",v,"Present?",ok)
}