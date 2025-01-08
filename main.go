package main

import (
	"fmt"
	"strings"
)

func main() {

	var conferenceName string = "Go Conference"
	const conferenceticket int = 50
	var remainTicket uint = 50
	//array
	// var bookings [50]string
	//slice
	var bookings []string

	greetUser(conferenceName, conferenceticket, remainTicket)

	for {

		isValidName, isValidEmail, isValidTicketNumber := validateUserInputs(firstName, lastName, email, userTicket, remainTicket)
		if isValidName && isValidEmail && isValidTicketNumber {
			remainTicket = remainTicket - userTicket

			//bookings[0] = firstName + " " + lastName

			bookings = append(bookings, firstName+" "+lastName)

			/* fmt.Printf("The whole array: %v\n", bookings)
			fmt.Printf("The first value: %v\n", bookings[0])
			fmt.Printf("Array Type: %T\n", bookings) */

			fmt.Printf("The whole slice: %v\n", bookings)

			fmt.Printf("The first value: %v\n", bookings[0])
			fmt.Printf("Array Type: %T\n", bookings)

			fmt.Printf("Array Length: %v\n", len(bookings))
			fmt.Printf("Thank you %v %v for booking %v tickets. You will receieve a confirmation email at %v.\n", firstName, lastName, userTicket, email)
			fmt.Printf("%v tickets remaining for %v.\n", remainTicket, conferenceName)

			fmt.Printf("These are all our booking : %v.\n", bookings)

			firstNames := getFirstName(bookings)

			fmt.Printf("The first names of booking are: %v.\n", firstNames)
			if remainTicket == 0 {
				fmt.Println("our Conforence is booked out, come back next year")
			}
		} else if userTicket == remainTicket {

		} else {

			if !isValidName {
				println("invalid name\n")
			}
			if !isValidEmail {
				fmt.Printf("Invalid email\n")
			}
			if !isValidTicketNumber {
				println("invlaid entry\n")
			}
			//fmt.Printf("We only have %v ticket remaining, so you cant book %v tickets.\n", remainTicket, userTicket)
			fmt.Printf("Your input data is invalid , try again.\n")
		}

	}

	city := "London"

	switch city {
	case "New York":
	case "Signapore", "HongKong":
	case "London", "Berlin":
	case "Mexico":
	default:
		fmt.Println("No Valid city selected")
	}
}

func greetUser(conferenceName string, conferenceticket int, remainTicket uint) {
	fmt.Printf("Welcome to %v booking application\n", conferenceName)
	fmt.Printf("We have %v ticket avaliable and %v ticket left\n", conferenceticket, remainTicket)
	fmt.Println("Thank you for booking a ticket for the go conference")
}

func getFirstName(bookings []string) []string {
	firstNames := []string{}

	for _, booking := range bookings {
		/* var names = strings.Fields(booking)
		var firstName = names[0]
		firstNames = append(firstNames, firstName) */

		var names = strings.Fields(booking)
		firstNames = append(firstNames, names[0])

	}

	return firstNames
}

func validateUserInputs(firstName string, lastName string, email string, userTicket uint, remainTicket uint) (bool, bool, bool) {
	isValidName := len(firstName) >= 2 && len(lastName) >= 2
	isValidEmail := strings.Contains(email, "@")
	isValidTicketNumber := userTicket >= 0 && userTicket <= remainTicket

	return isValidName, isValidEmail, isValidTicketNumber
}

func getUserInput() {}

/* func main() {
	const name, age = "Kim", 22
	fmt.Printf("%s is %d years old.\n", name, age)

	// It is conventional not to worry about any
	// error returned by Printf.

}
*/
