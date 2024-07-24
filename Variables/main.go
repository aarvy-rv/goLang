package main

import "fmt"

func main() {

	var name string = "mario" // one way to declare and initialize the variable
	var nameTwo = "Raj"       // second way where variable automatically infer the type according to the value assigned
	var nameThree string      // third way to just delare the variable to be used in future,by default for dtring type empty string is initialized
	fmt.Println(name, nameTwo, nameThree)

	name = "Qwerty"
	nameTwo = "kop"

	fmt.Println(name, nameTwo, nameThree)

	nameFour := "Vaibhav" // fourth way to use short hand notations(colon equal to )-> this initialization and delaration cannot be done outside the function

	fmt.Println(name, nameTwo, nameThree, nameFour)

}
