package main

import (
	"errors"
	"fmt"
	"time"
)

type user struct { // 1) We can also define it inside the class | 2) In Go, Casing of the first letter into upper case makes it available to the other packages.For instance, if user's U is capital then its available to ther package otherwise not.
	fistName  string
	lastName  string
	birthDate string
	createdAt time.Time
}

// Attaching methods to a struct, (u user) is known reciever argument and works on the concept of
// pass by value. To make it pass by reference we need pass the poiter to the struct in the reciever
// argumenr(u *user)
func (u *user) outputUserDetails() {
	fmt.Println(u.fistName, u.lastName, u.birthDate, u.createdAt)
}

func (u *user) clearUserName() {
	u.fistName = ""
	u.lastName = ""
}

//We can make a constructor to a struct to simplify struct creation, it is not a feature but
// more of a convention

func newUser(userFirstName, userLastName, userBirthDate string) (*user, error) {
	if userFirstName == "" || userLastName == "" || userBirthDate == "" {
		return nil, errors.New("FirstName,LastName or BirthDate is null")
	}
	return &user{
		fistName:  userFirstName,
		lastName:  userLastName,
		birthDate: userBirthDate,
		createdAt: time.Now(),
	}, nil
}
func main() {

	userFirstName := getUserData("Please enter your first name: ")
	userLastName := getUserData("Please enter your last name: ")
	userBirthDate := getUserData("Please enter your birthdate (MM/DD/YYYY): ")

	var appUser user

	appUser = user{
		fistName:  userFirstName,
		lastName:  userLastName,
		birthDate: userBirthDate,
		createdAt: time.Now(),
	}
	// If the order of the variables defined in the struct is same as while we instantiate a struct
	// variable then we can remove the key in the instantiation and just write the value
	// For eg:
	// appUser = user{
	// 		userFirstName,
	//		userLastName,
	//		userBirtDate,
	//		time.Now(),
	//     }
	// it becomes error prone also as if we incorrectly place the values then data would become incorrect

	// we can define empty struct : appUser= user{}

	/*outputUserDetails(&appUser) */
	appUser.outputUserDetails()
	appUser.clearUserName()
	appUser.outputUserDetails()

	var appUser2, err = newUser(userFirstName, userLastName, userBirthDate)
	if err != nil {
		fmt.Println(err)
		return
	}
	appUser2.outputUserDetails()

}

/*func outputUserDetails(u *user) {
	fmt.Println(u.fistName, u.lastName, u.birthDate, u.createdAt)
}*/ // We need not to derefernce the address for struct pointer like ((*u).firstName) because it automatically does that

func getUserData(promptText string) string {
	fmt.Print(promptText)
	var value string
	fmt.Scanln(&value) //Scanln moves to next command when hit enter
	return value
}
