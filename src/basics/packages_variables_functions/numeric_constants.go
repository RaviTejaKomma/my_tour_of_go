package main

import "fmt"

const (
	// Create a huge number by shifting a 1 bit left 100 places.
	// In other words, the binary number that is 1 followed by 100 zeroes.
	Big = 1 << 100
	// Shift it right again 99 places, so we end up with 1<<1, or 2.
	Small = Big >> 98
)

func needInt(x int) int {
	return 10*x + 1
}

func needFloat(x float64) float64{
	return  x * 0.1
}

func main(){
	fmt.Println(Small)
	// fmt.Println((Big) Throws error
	fmt.Println(float64(Big))
	fmt.Println(needInt(Small))
	fmt.Println(needFloat(Small))
	fmt.Println(needFloat(Big))
}