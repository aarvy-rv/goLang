package main

import (
	"fmt"
	"math"
)

func main() {

	const inflationRate float64 = 2.5 // constants always need to be assigneds

	// principal := 10000

	// principal, years := 10000,19

	// principal, years := 10000,"hello"

	var principal, years float64 = 10000, 19 // if we explicitly specify the type then all the comma seperated variables will be
	// of that type, we cannot specify different types for different comma sep. variable

	rate := 7.7

	var a string
	var b int = 2
	fmt.Println(a, b)
	// years := 19

	//returnAmount := float64(principal) * math.Pow((1+rate/100), float64(years))

	returnAmount := principal * math.Pow((1+rate/100), years)

	var realReturnValue = returnAmount / math.Pow(1+inflationRate/100, years)

	fmt.Println("Return Amount:", returnAmount)
	fmt.Println("Real Return Amount:", realReturnValue)
}
