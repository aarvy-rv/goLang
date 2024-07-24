package main

import (
	"fmt"
)

func main() {
	// By var keyword
	var arr = [3]int{22, 23, 24}             // array length is defined
	var arr1 = [...]int{1, 2, 3, 4, 5, 6, 7} // array length is inferred

	//By := assigment operator
	arr2a := [5]int{1, 2, 3, 4, 5}
	arr2b := [...]int{3, 4, 3}

	for i := 0; i < len(arr2a); i++ {
		fmt.Printf("i = %d\n", arr2a[i])
	}

	for i := 0; i < 5; i++ {
		for j := 0; j <= i; j++ {
			fmt.Print("*")
		}
		fmt.Println()
	}

	fmt.Println(arr)
	fmt.Println(arr1)
	fmt.Println(arr2a)
	fmt.Println(arr2b)

	prices := []float64{}
	latestPrices := append(prices, 7.8)
	fmt.Println(latestPrices)
	var i float64
	for i = 1; i < 9; i++ {
		prices = append(prices, i)
		fmt.Println(prices)
		var p *float64 = &prices[0]
		fmt.Println(p)

	}
	fmt.Println(prices)

	web := map[string]string{"Google": "www.google.com", "Amazon": "https://aws.com"}
	fmt.Println(web)

}
