package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

func getUserInput(text string) (float64, error) {
	fmt.Println(text)
	var input float64
	fmt.Scan(&input)

	if input <= 0 {
		return 1, errors.New("error!enter the valid input")
	}

	return input, nil
}

func writeToFile(profit float64, fileName string) {
	profitText := fmt.Sprint(profit)
	data := os.WriteFile(fileName, []byte(profitText), 0644)

	if data != nil {
		panic("file not found")
	}
}

func readFile(fileName string) {
	data, err := os.ReadFile(fileName)
	if err != nil {
		panic("File Not found")
	}

	dataString := string(data)
	value, e := strconv.ParseFloat(dataString, 64)
	if e != nil {
		panic("error! parsing of float")
	}
	fmt.Printf("Profit: %.3f", value)
}
func checkError(err error) {
	if err != nil {
		fmt.Println("Error! Wrong Input provided")
		panic("---------ERROR------------")
	}
}
func main() {
	var revenue, expenses, tax float64
	var err error
	revenue, err = getUserInput("Enter the revenue: ")
	checkError(err)
	expenses, err = getUserInput("Enter the expenses: ")
	checkError(err)
	tax, err = getUserInput("Enter tax rate: ")
	checkError(err)
	profit := revenue - expenses
	eat := profit - ((tax / 100) * revenue)

	writeToFile(profit, "ProfitBeforeTax.txt")
	writeToFile(eat, "ProfitAfterTax.txt")
	readFile("ProfitBeforeTax.txt")
	readFile("ProfitAfterTax.txt")
}
