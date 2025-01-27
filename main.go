package main

import (
	"fmt"
	"strings"
)

func main() {
	var name string = "Go Conference"
	const conferenceTickets uint = 50
	var remainingTickets uint = 50

	greetUsers(name)

	var bookings [50]string
	var bookingIndex int = 0

	for remainingTickets > 0 {
		var userName string
		var email string
		var userTickets uint

		fmt.Println("Enter your name:")
		fmt.Scan(&userName)

		fmt.Println("Enter your email:")
		fmt.Scan(&email)

		fmt.Println("Enter number of tickets:")
		fmt.Scan(&userTickets)

		if userTickets > remainingTickets {
			fmt.Printf("Sorry, we only have %v tickets remaining. Please try again.\n", remainingTickets)
			continue
		} else if userTickets <= 0 {
			fmt.Println("The number of tickets must be greater than 0. Please try again.")
			continue
		} else if len(userName) < 2 {
			fmt.Println("Name must be at least 2 characters long. Please try again.")
			continue
		} else if !strings.Contains(email, "@") {
			fmt.Println("Invalid email address. Please try again.")
			continue
		}

		remainingTickets -= userTickets
		bookings[bookingIndex] = userName
		bookingIndex++

		fmt.Printf("Thanks for booking %v tickets, %v! You will receive a confirmation email at %v.\n", userTickets, userName, email)
		fmt.Printf("Remaining tickets: %v\n", remainingTickets)

		if remainingTickets == 0 {
			fmt.Println("Sorry! No tickets are available. All tickets have been booked.")
			break
		}
	}

	fmt.Println("Here are all the bookings:")
	for i := 0; i < bookingIndex; i++ {
		fmt.Println(bookings[i])
	}
}

func greetUsers(name string) {
	fmt.Printf("Welcome to %v booking application!\n", name)
	fmt.Println("We have a total of 50 tickets, and they're going fast. Book your spot now!")
}
