package user

import (
	"errors"
	"fmt"
	"strings"
	"time"
	"unicode"
)

type User struct {
	firstName string
	lastName  string
	birthdate time.Time
	createdAt time.Time
}

type Admin struct {
	email string
	password string
	User
}

func (u *User) OutputUserDetails() {
	fmt.Printf("Name: %s %s | DOB: %s | CreatedAt: %s\n",
		u.firstName,
		u.lastName,
		u.birthdate.Format("01/02/2006"),
		u.createdAt.Format(time.RFC3339),
	)
}

func (u *User) ClearUserName() {
	u.firstName = ""
	u.lastName = ""
}

func isValidName(s string) bool {
	s = strings.TrimSpace(s)
	if s == "" {
		return false
	}
	for _, r := range s {
		if r == ' ' {
			continue
		}
		if !unicode.IsLetter(r) {
			return false
		}
	}
	return true
}

func NewAdmin(email, password string) Admin {
	return Admin{
		email: email,
		password: password,
		User: User{
			firstName: "ADMIN",
			lastName: "ADMIN",
			birthdate: time.Now() ,
			createdAt: time.Now(),
		},
	}
}

func newUser(firstName, lastName, birthdate string) (*User, error) {
	firstName = strings.TrimSpace(firstName)
	lastName = strings.TrimSpace(lastName)
	birthdate = strings.TrimSpace(birthdate)

	if !isValidName(firstName) {
		return nil, errors.New("first name must contain only letters")
	}
	if !isValidName(lastName) {
		return nil, errors.New("last name must contain only letters")
	}

	dob, err := time.Parse("01/02/2006", birthdate)
	if err != nil {
		return nil, errors.New("birthdate must be in MM/DD/YYYY format (e.g., 12/20/2000)")
	}

	now := time.Now()

	if dob.After(now) {
		return nil, errors.New("birthdate cannot be in the future")
	}

	if dob.Before(now.AddDate(-120, 0, 0)) {
		return nil, errors.New("birthdate looks too old (older than 120 years)")
	}

	return &User{
		firstName: firstName,
		lastName:  lastName,
		birthdate: dob,
		createdAt: now,
	}, nil
}

func NewUserFromInput() *User {
	for {
		firstName := getUserData("Please enter your first name: ")
		lastName := getUserData("Please enter your last name: ")
		birthdate := getUserData("Please enter your birthdate (MM/DD/YYYY): ")

		u, err := newUser(firstName, lastName, birthdate)
		if err != nil {
			fmt.Println("❌ Invalid input:", err)
			fmt.Println("Please try again.")
			continue
		}
		return u
	}
}

func getUserData(promptText string) string {
	fmt.Print(promptText)
	var value string
	fmt.Scan(&value)
	return value
}
