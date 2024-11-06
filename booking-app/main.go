package main

import (
	"fmt"
	"strings"
)

// Use Println with variables and constants
func main() {
	conferenceName := "Go Conference"
	// conferenceTickets is a constant instead of a variable because it does not change throughout the application
	const conferenceTickets = 50
	// Keep track of ticket count to know how many are still available
	// Explicitely define uint as the data type so that it does not accept negative numbers
	var remainingTickets uint = 50

	fmt.Println("Welcome to the", conferenceName, "booking application!")
	fmt.Println("We have a total of", conferenceTickets, "tickets and", remainingTickets, "are still available.")
	fmt.Println("Get your tickets here to attend.")

	// Printf is used for printing formatted data (%v, %s, etc.)
	// Use \n at the end to escape to a new line (Println automatically adds a new line, but Printf does not)
	// Reference the variable to be inserted at the %v spot at the end of the quotation
	fmt.Printf("Welcome to %v booking application!\n", conferenceName)
	fmt.Printf("We have a total of %v tickets and %v are still available.\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend.")

	// Use %T to print the data type
	fmt.Printf("conferenceTickets is %T, remainingTickets is %T, conferenceName is %T\n", conferenceTickets, remainingTickets, conferenceName)

	var firstName string
	var lastName string
	var email string
	var userTickets uint

	// Ask for user input
	fmt.Println("Enter your first name: ")
	fmt.Scan(&firstName)

	fmt.Println("Enter your last name: ")
	fmt.Scan(&lastName)

	fmt.Println("Enter your email: ")
	fmt.Scan(&email)

	fmt.Println("Enter the number of tickets you want to book: ")
	fmt.Scan(&userTickets)

	remainingTickets = remainingTickets - userTickets

	// Store a list of the inputted user names via an array of string type
	var bookings = [50]string{}
	bookings[0] = firstName + " " + lastName

	fmt.Printf("The whole array: %v\n", bookings)
	fmt.Printf("The first value: %v\n", bookings[0])
	fmt.Printf("The type of the array: %T\n", bookings)
	fmt.Printf("The size of the array: %v\n", len(bookings))

	// Example of how to do bookings as a slice
	var bookings2 []string
	bookings2 = append(bookings2, firstName+" "+lastName)

	fmt.Printf("The whole slice: %v\n", bookings2)
	fmt.Printf("The first value: %v\n", bookings2[0])
	fmt.Printf("The type of the slice: %T\n", bookings2)
	fmt.Printf("The size of the slice: %v\n", len(bookings2))

	fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation email at %v\n", firstName, lastName, userTickets, email)
	fmt.Printf("There are %v tickets remaining for the %v\n", remainingTickets, conferenceName)

	fmt.Printf("These are all our bookings: %v\n", bookings2)

	forLoop()
}

// Use for loop to allow booking over and over again
func forLoop() {
	conferenceName := "Go Conference"
	const conferenceTickets = 50
	var remainingTickets uint = 50

	// Declare the bookings2 slice
	var bookings2 []string

	fmt.Printf("Welcome to %v booking application!\n", conferenceName)
	fmt.Printf("We have a total of %v tickets and %v are still available.\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend.")

	for {
		var firstName string
		var lastName string
		var email string
		var userTickets uint

		// Ask for user input
		fmt.Println("Enter your first name: ")
		fmt.Scan(&firstName)

		fmt.Println("Enter your last name: ")
		fmt.Scan(&lastName)

		fmt.Println("Enter your email: ")
		fmt.Scan(&email)

		fmt.Println("Enter the number of tickets you want to book: ")
		fmt.Scan(&userTickets)

		remainingTickets = remainingTickets - userTickets

		// Example of how to do bookings as a slice
		bookings2 = append(bookings2, firstName+" "+lastName)

		fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation email at %v\n", firstName, lastName, userTickets, email)
		fmt.Printf("There are %v tickets remaining for the %v\n", remainingTickets, conferenceName)

		// Declare slice to only print firstName for privacy
		firstNames := []string{}
		for _, booking := range bookings2 {
			var names = strings.Fields(booking)
			firstNames = append(firstNames, names[0])
		}

		fmt.Printf("These first names of bookings are: %v\n", firstNames)

		// Check if there are any tickets left
		if remainingTickets == 0 {
			// End the program (for loop)
			fmt.Println("Our conference is booked out. Please come back next year.")
			break
		}
	}
}
