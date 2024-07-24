package main

import "fmt"

//type keyword can be used to alias
type str string

func (text str) log() {
	fmt.Println(text)
}
func main() {
	var name str = "Raj"
	name.log()
}
