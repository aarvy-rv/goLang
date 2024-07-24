package main

import "fmt"

func main(){
	var firstName,lastName,birthDate string 
	GetUserInput(&firstName,"Enter the first name: ")
	GetUserInput(&lastName,"Enter the Last Name")
	GetUserInput(&birthDate,"Enter the birth date")
}

func GetUserInput(input *string, inputText string){
	fmt.Print(inputText)
	fmt.Scan(input)
}