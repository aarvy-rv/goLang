package main

import (
	"fmt"
	"math"
)

const inflationRate = 2.2

func main() {

	fmt.Print(`Welcome Guys!
Type Valid Values
`) // using `` instead of "" outputs the text in the similar manner as written inside the parameter as it is,
	//backslash charaters too are ignored in it

	var investment float64

	investment = input("Enter Investment: ", &investment)
	fmt.Println("Investment: ", investment)
	//fmt.Scan(&investment)

	var rate float64
	fmt.Print("Rate: ")
	fmt.Scan(&rate)

	var years float64
	fmt.Print("Years: ")
	fmt.Scan(&years)

	futureValues, realFutureValues := calculateValues(investment, rate, years)

	text := fmt.Sprintf("Future Value: %.2f\nrealFutureValues: %.2f", futureValues, realFutureValues)

	outputText(text)

	a, b := calculate(89, 67)
	var p, q float64 = calculate(12, 12)
	fmt.Println(p, q)
	fmt.Printf("Sum of 89 and 67: %f\nProduct of 89 and 67:%f", a, b)
}

func outputText(text string) {
	fmt.Println(text)
}

func input(text string, variable *float64) float64 {
	fmt.Print(text)
	fmt.Scan(variable)
	return *variable
}

func calculateValues(investment, returnRate, years float64) (float64, float64) { //if parameter values are of same type then we can separate variables with comma and lastly mention datatype

	fv := investment * math.Pow(1+returnRate/100, years)
	rfv := fv / math.Pow(1+inflationRate/100, years)

	return fv, rfv

}

/* funct calculateValues(investment float64,returnRate float64,years float64) (float64,float64){
	fv := investment * math.Pow(1+returnRate/100,years)
	rfv := fv / math.Pow(1+inflationRate/100,years)

	return fv,rfv
}
*/

// Another way to return value from function
func calculate(a, b float64) (x float64, y float64) { // Values returned
	x = a + b
	y = a * b

	return
}
