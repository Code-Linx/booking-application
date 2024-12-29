package main

import (
	"fmt"
)

func main() {

	var conferenceName string = "Go Conference"
	const conferenceticket int = 50
	var remainTicket uint = 50
	var bookings [50]string

	fmt.Printf("Welcome to %v  booking application\n", conferenceName)
	fmt.Printf("We have %v ticket avaliable and %v ticket left\n", conferenceticket, remainTicket)
	fmt.Println("Thank you for booking a ticket for the go conference")

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

	remainTicket = remainTicket - userTicket

	bookings[0] = firstName + " " + lastName
	fmt.Printf("The whole array: %v\n", bookings)
	fmt.Printf("The first value: %v\n", bookings[0])
	fmt.Printf("Array Type: %T\n", bookings)
	fmt.Printf("Array Length: %v\n", len(bookings))
	fmt.Printf("Thank you %v %v for booking %v tickets. You will receieve a confirmation email at %v.\n", firstName, lastName, userTicket, email)
	fmt.Printf("%v tickets remaining for %v.\n", remainTicket, conferenceName)
}

/* func main() {
	const name, age = "Kim", 22
	fmt.Printf("%s is %d years old.\n", name, age)

	// It is conventional not to worry about any
	// error returned by Printf.

}
*/
