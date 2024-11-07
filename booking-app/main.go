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

	// Can call another function to greet the users
	greetUsers(conferenceName, conferenceTickets, remainingTickets)

	// Or can have the greeting like this
	fmt.Println("Welcome to the", conferenceName, "booking application!")
	fmt.Println("We have a total of", conferenceTickets, "tickets and", remainingTickets, "are still available.")
	fmt.Println("Get your tickets here to attend.")

	// Printf is used for printing formatted data (%v, %s, etc.)
	// Use \n at the end to escape to a new line (Println automatically adds a new line, but Printf does not)
	// Reference the variable to be inserted at the %v spot at the end of the quotation
	// Or can have the greeting like this
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
	var remainingTickets uint = 50

	// Declare the bookings2 slice
	var bookings2 []string

	// Move user input code to its own function
	for remainingTickets > 0 {
		firstName, lastName, email, userTickets := getUserInput()

		// Move validation to its own function
		isValidName, isValidEmail, isValidUserTickets := validateUserInput(firstName, lastName, email, userTickets, remainingTickets)

		if isValidName && isValidEmail && isValidUserTickets {
			// Move code for booking to its own function
			bookTicket(remainingTickets, userTickets, bookings2, firstName, lastName, email, conferenceName)

			// Move firstNames section to its own function
			// Call function to print firstNames
			firstNames := getFirstNames(bookings2)
			fmt.Printf("The first names of bookings are: %v\n", firstNames)

			// Check if there are any tickets left
			if remainingTickets == 0 {
				// End the program (for loop)
				fmt.Println("Our conference is booked out. Please come back next year.")
				break
			}

		} else if userTickets == remainingTickets {
			fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation email at %v\n", firstName, lastName, userTickets, email)
			fmt.Printf("There are 0 tickets remaining for the %v\n", conferenceName)
			break

		} else if userTickets > remainingTickets {
			fmt.Printf("I'm sorry, you can't book %v tickets because we only have %v tickets remaining. Please try again.\n", userTickets, remainingTickets)
			continue

		} else {
			if !isValidName {
				fmt.Println("Error: The first name or last name entered is too short. Please try again.")
			}
			if !isValidEmail {
				fmt.Println("Error: The email address entered does not contain the @ sign. Please try again.")
			}
			if !isValidUserTickets {
				fmt.Println("The number of tickets entered is invalid. Please try again.")
			}
		}
	}
}

func greetUsers(conferenceName string, conferenceTickets int, remainingTickets uint) {
	fmt.Printf("Welcome to our %v booking application!\n", conferenceName)
	fmt.Printf("We have a total of %v tickets and %v are still available.\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend.")
}

func getFirstNames(bookings2 []string) []string {
	// Declare slice to only print firstName for privacy
	firstNames := []string{}
	for _, booking := range bookings2 {
		var names = strings.Fields(booking)
		firstNames = append(firstNames, names[0])
	}
	// Move the printf firstNames line to the main function and just return the firstNames here so that the main function (using func forLoop for this) has them
	return firstNames // The output of this will get printed in a prior function
	// You need to tell the function what data type the return tool is returning, so add []string as a parameter of the printFirstNames function, so it knows it is returning a list of strings
}

func validateUserInput(firstName string, lastName string, email string, userTickets uint, remainingTickets uint) (bool, bool, bool) {
	var isValidName = len(firstName) >= 2 && len(lastName) >= 2
	var isValidEmail = strings.Contains(email, "@")
	var isValidUserTickets = userTickets > 0 && userTickets <= remainingTickets
	return isValidName, isValidEmail, isValidUserTickets
}

func getUserInput() (string, string, string, uint) {
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

	return firstName, lastName, email, userTickets
}

func bookTicket(remainingTickets uint, userTickets uint, bookings2 []string, firstName string, lastName string, email string, conferenceName string) {
	remainingTickets = remainingTickets - userTickets

	// Example of how to do bookings as a slice
	bookings2 = append(bookings2, firstName+" "+lastName)

	fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation email at %v\n", firstName, lastName, userTickets, email)
	fmt.Printf("There are %v tickets remaining for the %v\n", remainingTickets, conferenceName)
}
