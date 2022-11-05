package main

import "fmt"

func main() {
	var conferenceName = "Go Conference"
	const conferenceTickets uint = 50
	var remainingTickets uint = 50
	var bookings []string

	fmt.Printf("Welcome to %v booking application\n", conferenceName)
	fmt.Printf("We have total of %v tickets and %v are still available\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend")

	for {
		var firstName string
		var lastName string
		var userTickets uint

		// ask user for their name
		fmt.Println("Please enter your first name")
		fmt.Scan(&firstName)

		fmt.Println("Please enter your last name")
		fmt.Scan(&lastName)

		// ask user for number of tickets
		fmt.Println("Please enter number of tickets")
		fmt.Scan(&userTickets)

		// update remaining tickets
		remainingTickets -= userTickets
		bookings = append(bookings, firstName+" "+lastName)

		fmt.Printf("Thank you %v for booking %v tickets\n", firstName, userTickets)
		fmt.Printf("We have %v tickets left\n", remainingTickets)
		fmt.Printf("Bookings so far: %v \n", bookings)
	}
}
