package main

import (
	"fmt"
)

var balance float64 = 100000

func main() {
	fmt.Println("Welcome to the Go Bank!")

	for {
		var choice int
		showOptions()
		fmt.Scan(&choice)

		switch choice {
		case 1:
			displayBalance()
		case 2:
			var amount float64
			fmt.Print("Enter the amount to withdraw: ")
			fmt.Scan(&amount)

			if amount <= 0 {
				fmt.Println("Please Enter amount above 0")
				continue
			}
			if amount > balance {
				fmt.Printf("Insufficient balance to withdraw %f amount! Please Enter valid amount\n", amount)
				continue
			}
			withdrawAmount(amount)

		case 3:
			fmt.Print("Enter the amount to deposit: ")
			var amount float64
			fmt.Scan(&amount)
			if amount <= 0 {
				fmt.Println("Please Entet amount above 0")
				continue
			}
			depositAmount(amount)
		case 4:
			fmt.Println("Thank You!!")
			fmt.Println("Have a good day!Thank you for choosing Go Bank\nPress Enter to exit!")
			var e string
			fmt.Scan(&e)
			return
		default:
			fmt.Println("Invalid Choice!")
		} //break keyword inside the switch statement brings the control outside the switch block instead of nearest loo
		//So,to exit one can use "return" keyword which exit control from the app
	}

}
func showOptions() {
	fmt.Println("Enter 1 to check balance")
	fmt.Println("Enter 2 to withdraw")
	fmt.Println("Enter 3 to deposit")
	fmt.Println("Enter 4 to exit")
	fmt.Println("Your choice: ")
}

func displayBalance() {
	fmt.Printf("Your Account Balance: %f\n", balance)
}

func depositAmount(amount float64) {
	balance += amount
	fmt.Printf("%f is deposited in your account\n", amount)
	displayBalance()
}

func withdrawAmount(amount float64) {
	balance -= amount
	displayBalance()
}
