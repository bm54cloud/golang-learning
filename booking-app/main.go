package main

import "fmt"

func main() {
	var conferenceName = "Go Conference"
	// conferenceTickets is a constant instead of a variable because it does not change throughout the application
	const conferenceTickets = 50
	// Keep track of ticket count to know how many are still available
	var remainingTickets = 50

	fmt.Println("Welcome to the", conferenceName, "booking application!")
	fmt.Println("We have a total of", conferenceTickets, "tickets and", remainingTickets, "are still available.")
	fmt.Println("Get your tickets here to attend")

	// Printf is used for printing formatted data
	// Use \n at the end to escape to a new line
	// Reference the data to be inserted at the %v spot at the end of the quotation
	fmt.Printf("Welcome to %v booking application!\n", conferenceName)



}
