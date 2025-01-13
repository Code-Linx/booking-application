package main

import (
	"fmt"
	"go-booking-app/helper"
	"time"
)

var conferenceName string = "Go Conference"

const conferenceticket int = 50

var remainTicket uint = 50

// array
// var bookings [50]string
// slice
// var bookings []string
var bookings = make([]UserData, 0)

type UserData struct {
	firstName       string
	lastName        string
	email           string
	NumbersOfTicket uint
}

func main() {

	greetUser()

	for {

		firstName, lastName, email, userTicket := getUserInput()

		isValidName, isValidEmail, isValidTicketNumber := helper.ValidateUserInputs(firstName, lastName, email, userTicket, remainTicket)
		if isValidName && isValidEmail && isValidTicketNumber {
			bookTickets(userTicket, firstName, lastName, email)
			go sendTicket(userTicket, firstName, lastName, email)

			firstNames := getFirstName()

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

	/* city := "London"
	switch city {
	case "New York":
	case "Signapore", "HongKong":
	case "London", "Berlin":
	case "Mexico":
	default:
		fmt.Println("No Valid city selected")
	} */
}

func greetUser() {
	fmt.Printf("Welcome to %v booking application\n", conferenceName)
	fmt.Printf("We have %v ticket avaliable and %v ticket left\n", conferenceticket, remainTicket)
	fmt.Println("Thank you for booking a ticket for the go conference")
}

func getFirstName() []string {
	firstNames := []string{}

	for _, booking := range bookings {
		/* var names = strings.Fields(booking)
		var firstName = names[0]
		firstNames = append(firstNames, firstName) */

		//var names = strings.Fields(booking)
		firstNames = append(firstNames, booking.firstName)

	}

	return firstNames
}

func getUserInput() (string, string, string, uint) {
	var firstName string
	var lastName string
	var email string
	var userTicket uint

	fmt.Println("Enter Your First Name:")
	fmt.Scan(&firstName)

	fmt.Println("Enter Your Last Name:")
	fmt.Scan(&lastName)

	fmt.Println("Enter Your Email address:")
	fmt.Scan(&email)

	fmt.Println("Enter Numbers of Ticket:")
	fmt.Scan(&userTicket)

	return firstName, lastName, email, userTicket
}

func bookTickets(userTicket uint, firstName string, lastName string, email string) {
	remainTicket = remainTicket - userTicket

	var userData = UserData{
		firstName:       firstName,
		lastName:        lastName,
		email:           email,
		NumbersOfTicket: userTicket,
	}

	/* userData["firstName"] = firstName
	userData["lastName"] = lastName
	userData["email"] = email
	userData["NumbersOfTicket"] = strconv.FormatUint(uint64(userTicket), 10) */
	//bookings[0] = firstName + " " + lastName

	bookings = append(bookings, userData)
	fmt.Printf("List of booking %v.\n", bookings)

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
}

func sendTicket(userTicket uint, firstName string, lastName string, email string) {
	time.Sleep(10 * time.Second)
	var ticket = fmt.Sprintf("%v tickets for %v %v.\n", userTicket, firstName, lastName)
	println(("######################"))
	fmt.Printf("Sending ticket :\n %v \nto email address  %v.\n", ticket, email)
	println(("######################"))
}

/* func main() {
	const name, age = "Kim", 22
	fmt.Printf("%s is %d years old.\n", name, age)

	// It is conventional not to worry about any
	// error returned by Printf.

}
*/
