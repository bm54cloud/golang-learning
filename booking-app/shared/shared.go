// These are helper or shared functions of the main package
// Let's change it to be its own package

// package main
package shared

import "strings"

// Variables and functions defined outside any function can be accessed in all other files within the same package
func ValidateUserInput(firstName string, lastName string, email string, userTickets uint, remainingTickets uint) (bool, bool, bool) {
	var isValidName = len(firstName) >= 2 && len(lastName) >= 2
	var isValidEmail = strings.Contains(email, "@")
	var isValidUserTickets = userTickets > 0 && userTickets <= remainingTickets
	return isValidName, isValidEmail, isValidUserTickets
}
