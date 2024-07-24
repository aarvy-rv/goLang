package main

import (
	"fmt"

	user "learning.com/go/17user"
)

func main() {
	fmt.Println("Hello")

	var appUser *user.User
	appUser, err := user.New("Raj", "Vaibhav", "12/31/1997")
	if err != nil {
		fmt.Println(err)
		return
	}
	appUser.OutputUserDetails()

	admin := user.NewAdmin("test@example.com", "test123")
	admin.OutputUserDetails()

}
