package main

import (
	"booking-app/util"
	"fmt"
)

func main() {
	appName := "My Booking App"
	const initialTickets = 50
	availableTickets := initialTickets

	fmt.Println("Welcome to", appName)
	fmt.Printf("There are currently %d out of %d tickets still available.\n", availableTickets, initialTickets)

	var bookings = make([]util.UserData, 0)

	for availableTickets > 0 {

		var userData, err = util.HandleInput(availableTickets)
		if err != nil {
			fmt.Println("Received error", err, "Please retry.")
			continue
		}
		bookings = append(bookings, userData)
		availableTickets -= userData.NumTickets

		fmt.Printf("Thank you %s %s for booking %d tickets. There are %d tickets remaining. You will receive your confirmation email at %s.\n",
			userData.FirstName, userData.LastName, userData.NumTickets, availableTickets, userData.Email)
	}

	for i, booking := range bookings {
		fmt.Println("Booking", i, "is", booking)
	}
}
