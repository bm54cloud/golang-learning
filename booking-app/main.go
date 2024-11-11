package main

import (
	"booking-app/shared"
	"fmt"
	"strconv"
	"time"
)

// Package level variables
var conferenceName = "Go Conference"

const conferenceTickets = 50   // conferenceTickets is a constant instead of a variable because it does not change throughout the application
var remainingTickets uint = 50 // Explicitely define uint as the data type so that it does not accept negative numbers
var bookings = [50]string{}

// Change bookings list to a map instead of a list of strings
// var bookings2 []string
// Change bookings2 to an empty slice of maps and define the starting size (0)
var bookings2 = make([]map[string]string, 0)

// Use Println with variables and constants
func main() {
	// Move variables to package level so they can be shared
	// Move greet users code to its own function and call itf
	greetUsers()

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
	bookings[0] = firstName + " " + lastName

	fmt.Printf("The whole array: %v\n", bookings)
	fmt.Printf("The first value: %v\n", bookings[0])
	fmt.Printf("The type of the array: %T\n", bookings)
	fmt.Printf("The size of the array: %v\n", len(bookings))

	var userData = make(map[string]string)
	userData["firstname"] = firstName
	userData["lastName"] = lastName
	userData["email"] = email
	userData["numberOfTickets"] = strconv.FormatUint(uint64(userTickets), 10)

	// Example of how to do bookings as a slice
	bookings2 = append(bookings2, userData)

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
	// Move user input code to its own function
	for remainingTickets > 0 {
		firstName, lastName, email, userTickets := getUserInput()

		// Move validation to its own function
		isValidName, isValidEmail, isValidUserTickets := shared.ValidateUserInput(firstName, lastName, email, userTickets, remainingTickets)

		if isValidName && isValidEmail && isValidUserTickets {
			// Move code for booking to its own function
			bookTicket(userTickets, firstName, lastName, email)
			go sendTicket(userTickets, firstName, lastName, email)

			// Move firstNames section to its own function
			// Call function to print firstNames
			firstNames := getFirstNames()
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

func greetUsers() {
	fmt.Printf("Welcome to our %v booking application!\n", conferenceName)
	fmt.Printf("We have a total of %v tickets and %v are still available.\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend.")
}

func getFirstNames() []string {
	// Declare slice to only print firstName for privacy
	firstNames := []string{}
	for _, booking := range bookings2 {
		firstNames = append(firstNames, booking["firstName"])
	}
	// Move the printf firstNames line to the main function and just return the firstNames here so that the main function (using func forLoop for this) has them
	return firstNames // The output of this will get printed in a prior function
	// You need to tell the function what data type the return tool is returning, so add []string as a parameter of the printFirstNames function, so it knows it is returning a list of strings
}

// Move validateUserInput function to shared.go

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

func bookTicket(userTickets uint, firstName string, lastName string, email string) {
	remainingTickets = remainingTickets - userTickets

	// Example of how to do bookings as a slice

	// Creat a map for a user
	// First create a variable of an empty map
	var userData = make(map[string]string)
	userData["firstname"] = firstName
	userData["lastName"] = lastName
	userData["email"] = email
	userData["numberOfTickets"] = strconv.FormatUint(uint64(userTickets), 10)

	// Add the map to the bookins2 list
	// Make the bookings2 variable into a slice of the map
	// Copy the data type of the map (map[string]string)
	bookings2 = append(bookings2, userData)
	fmt.Printf("List of bookings is %v\n", bookings2)

	fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation email at %v\n", firstName, lastName, userTickets, email)
	fmt.Printf("There are %v tickets remaining for the %v\n", remainingTickets, conferenceName)
}

func sendTicket(userTickets uint, firstName string, lastName string, email string) {
	fmt.Println("Processing...")
	time.Sleep(10 * time.Second) // Sleep for 10 seconds
	var ticket = fmt.Sprintf("%v tickets for %v %v\n", userTickets, firstName, lastName)
	fmt.Println("#######") // Adds a visual divider
	fmt.Printf("Sending ticket:\n %v \nto email address %v\n", ticket, email)
	fmt.Println("#######")
}
