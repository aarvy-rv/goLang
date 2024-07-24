package main

import (
	"fmt"
	"math"
)

func main() {
	var investment float64
	var roi float64
	var years float64

	fmt.Print("Enter the investment amount: ")
	fmt.Scan(&investment)

	fmt.Print("Enter the rate of interest: ")
	fmt.Scan(&roi)

	fmt.Print("Enter the years of investment: ")
	fmt.Scan(&years)

	result := investment * math.Pow(1+roi/100, years)
	fmt.Println("Return of investment:", result)

	fmt.Print("Enter the inflation rate: ")
	var inflationRate float64
	fmt.Scan(&inflationRate)

	resultAfterInflation := result / math.Pow(1+inflationRate/100, years)
	fmt.Println("Return after taking inflation into consideration", resultAfterInflation)

	fmt.Printf("Output 1: %v \n Output 2: %v \n", result, resultAfterInflation)

	fmt.Printf("RA: %.2f\nIRA: %.0f", result, resultAfterInflation)
}
