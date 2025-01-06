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

	fmt.Printf("Welcome to %v  booking application\n", conferenceName)
	fmt.Printf("We have %v ticket avaliable and %v ticket left\n", conferenceticket, remainTicket)
	fmt.Println("Thank you for booking a ticket for the go conference")

	for remainTicket > 0 && len(bookings) < 50 {
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

		if userTicket < remainTicket {
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

			firstNames := []string{}

			for _, booking := range bookings {
				/* var names = strings.Fields(booking)
				var firstName = names[0]
				firstNames = append(firstNames, firstName) */

				var names = strings.Fields(booking)
				firstNames = append(firstNames, names[0])

			}
			fmt.Printf("The first names of booking are: %v.\n", firstNames)
			fmt.Printf("These are all our booking : %v.\n", bookings)
			if remainTicket == 0 {
				fmt.Println("our Conforence is booked out, come back next year")
			}
		} else if userTicket == remainTicket {

		} else {
			fmt.Printf("We only have %v ticket remaining, so you cant book %v tickets.\n", remainTicket, userTicket)

		}

	}
}

/* func main() {
	const name, age = "Kim", 22
	fmt.Printf("%s is %d years old.\n", name, age)

	// It is conventional not to worry about any
	// error returned by Printf.

}
*/
