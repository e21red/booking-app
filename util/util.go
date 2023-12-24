package util

import (
	"errors"
	"fmt"
	"strings"
)

type UserData struct {
	FirstName  string
	LastName   string
	Email      string
	NumTickets int
}

var InputError = errors.New("Invalid Input")

func ValidString(str string) bool {
	if len(str) < 2 {
		return false
	}
	return true
}

func ValidateEmail(str string) bool {
	return ValidString(str) && strings.Contains(str, "@")
}

func HandleInput(limit int) (UserData, error) {
	var retVal UserData

	var firstName string
	var lastName string
	var email string
	var numTickets int
	// asks user for input
	fmt.Println("Input first name:")
	fmt.Scan(&firstName)
	if !ValidString(firstName) {
		fmt.Println("Invalid name", firstName)
		return retVal, InputError
	}
	fmt.Println("Input last name:")
	fmt.Scan(&lastName)
	if !ValidString(lastName) {
		fmt.Println("Invalid name", lastName)
		return retVal, InputError
	}
	fmt.Println("Input your email:")
	fmt.Scan(&email)
	if !ValidateEmail(email) {
		fmt.Println("Invalid email", email)
		return retVal, InputError
	}
	fmt.Println("Input num tickets:")
	fmt.Scan(&numTickets)
	if numTickets < 0 || numTickets > limit {
		fmt.Println("Invalid number of tickets", numTickets)
		return retVal, InputError
	}
	retVal = UserData{
		FirstName:  firstName,
		LastName:   lastName,
		Email:      email,
		NumTickets: numTickets,
	}

	return retVal, nil
}
