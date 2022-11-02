package main

import "fmt"

func main() {
	var conferenceName = "Go Conference"
	const conferenceTickets uint = 50
	var remainingTickets uint = 50

	fmt.Printf("Welcome to %v booking application\n", conferenceName)
	fmt.Printf("We have total of %v tickets and %v are still available\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend")

	var userName string
	var userTickets int
	// ask user for their name
	fmt.Println("Please enter your name")
	fmt.Scan(&userName)
	// ask user for number of tickets
	fmt.Println("Please enter number of tickets")
	fmt.Scan(&userTickets)
	// check if user has entered valid number of tickets
	if userTickets > int(remainingTickets) {
		fmt.Println("Sorry, we don't have that many tickets left")
	} else {
		// update remaining tickets
		remainingTickets -= uint(userTickets)
		fmt.Printf("Thank you %v for booking %v tickets\n", userName, userTickets)
		fmt.Printf("We have %v tickets left\n", remainingTickets)
	}

	// userName = "Toma"
	// userTickets = 2
	fmt.Printf("User %v booked %v tickets\n", userName, userTickets)

}
