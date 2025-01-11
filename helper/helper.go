package helper

import "strings"

func ValidateUserInputs(firstName string, lastName string, email string, userTicket uint, remainTicket uint) (bool, bool, bool) {
	isValidName := len(firstName) >= 2 && len(lastName) >= 2
	isValidEmail := strings.Contains(email, "@")
	isValidTicketNumber := userTicket >= 0 && userTicket <= remainTicket

	return isValidName, isValidEmail, isValidTicketNumber
}
