package user

import (
	"errors"
	"fmt"
	"time"
)

// To make member variables of struct available outside the current package,
// we need to make the first letter of that varialble to be capital
type User struct {
	fistName  string
	lastName  string
	birthDate string
	createdAt time.Time
}

//embedding struct : similar to inheritance

type Admin struct {
	email    string
	password string
	User
}

func NewAdmin(email, password string) *Admin {
	return &Admin{
		email:    email,
		password: password,
		User: User{
			fistName:  "ADMIN",
			lastName:  "ADMIN",
			birthDate: "----",
			createdAt: time.Now(),
		}}
}

func New(userFirstName, userLastName, userBirthDate string) (*User, error) {
	if userFirstName == "" || userLastName == "" || userBirthDate == "" {
		return nil, errors.New("FirstName,LastName or BirthDate is null")
	}
	return &User{
		fistName:  userFirstName,
		lastName:  userLastName,
		birthDate: userBirthDate,
		createdAt: time.Now(),
	}, nil
}

func (u *User) OutputUserDetails() {
	fmt.Println(u.fistName, u.lastName, u.birthDate, u.createdAt)
}

func (u *User) ClearUserName() {
	u.fistName = ""
	u.lastName = ""
}
