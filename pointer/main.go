package main

import "fmt"

func main() {

	num1 := 2
	num2 := 5

	var str string = "Hello! World"
	var pointerVariable *string = &str //pointer to a string
	fmt.Println("Address:", pointerVariable)

	fmt.Printf("Before Swapping:\nnum1: %d,num2: %d", num1, num2)
	swap(&num1, &num2)
	fmt.Printf("\nAfter Swapping:\nnum1: %d,num2: %d", num1, num2)
}
func swap(a, b *int) {
	var temp int = *a
	*a = *b
	*b = temp
}
