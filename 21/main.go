package main

import (
	"errors"
	"fmt"
)

type transformFn func(*int) int

func main() {
	numbers := []int{1, 2, 3, 4} //slice

	var dNumbers []int
	for i := 0; i < len(numbers); i++ {
		dNumbers = append(dNumbers, numbers[i]*2)
	}
	fmt.Println(dNumbers)
	doubleNumber(&numbers)

	//Passing function as a value parameter in a function argument
	transformNumber(&numbers, double)
	transformNumber(&numbers, triple)

	//Anonymous Function
	transformNumber(&numbers, func(val *int) int {
		return *val * 4
	})

	integers := []int{23, 24, 25, 26}

	func1, err := getTransformFn("double")
	if err != nil {
		fmt.Println(err)

	} else {
		transformNumber(&integers, func1)
	}

	func2, err := getTransformFn("abc")
	if err != nil {
		fmt.Println(err)
	} else {
		transformNumber(&integers, func2)
	}

	func3, err := getTransformFn("triple")
	if err != nil {
		fmt.Println(err)
	} else {
		transformNumber(&integers, func3)
	}

}

func doubleNumber(input *[]int) {

	var dNumbers []int
	for _, val := range *input {
		dNumbers = append(dNumbers, val*2)
	}
	fmt.Println(dNumbers)
}

func transformNumber(numbers *[]int, transform transformFn) {

	var transformedSlice []int

	for _, val := range *numbers {
		transformedSlice = append(transformedSlice, transform(&val))
	}

	fmt.Println("Output: ", transformedSlice)
}

func getTransformFn(functionName string) (transformFn, error) {

	switch functionName {
	case "double":
		return double, nil
	case "triple":
		return triple, nil
	default:
		return nil, errors.New("invalid function name")
	}

}

func double(num *int) int {
	return *num * 2
}

func triple(num *int) int {
	return *num * 3
}
