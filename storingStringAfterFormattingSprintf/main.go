package main

import (
	"fmt"
)

func main() {

	var revenue, expense, taxRate float64
	fmt.Print("Enter the revenue: ")
	fmt.Scan(&revenue)

	fmt.Print("Enter the expenses: ")
	fmt.Scan(&expense)

	profit := revenue - expense

	fmt.Print("Enter tax rate: ")
	fmt.Scan(&taxRate)

	eat := profit - (1 - taxRate/100)

	formattedText1 := fmt.Sprintf("Profit: %.2f\nEarnings after tax: %.2f\n", profit, eat)

	fmt.Print(formattedText1)

}
