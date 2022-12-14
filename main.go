package main

import (
	"fmt"
	"strings"
)

func main() {
	conferenceName := "GO Conference"
	const confereceTickets int = 50
	var remainingTickets uint = 50

	fmt.Printf("Welcome to %v booking application\n", conferenceName)
	fmt.Printf("We have a total of %v tickets and %v are still available.\n", confereceTickets, remainingTickets)
	fmt.Printf("Get your tickets to attend.\n")
	bookings := []string{}

	for {
		var firstName string
		var lastName string
		var email string
		var userTickets uint

		fmt.Println("Enter your first name: ")
		fmt.Scan(&firstName)
		fmt.Println("Enter your last name: ")
		fmt.Scan(&lastName)
		fmt.Println("Enter your email: ")
		fmt.Scan(&email)
		fmt.Println("Enter number of tickets: ")
		fmt.Scan(&userTickets)

		if userTickets < remainingTickets {
			remainingTickets = remainingTickets - userTickets

			bookings = append(bookings, firstName+" "+lastName)

			fmt.Printf("Thank you %v %v for booking %v tickets. You'll receive a confirmation email at %v\n", firstName, lastName, userTickets, email)
			fmt.Printf("%v tickets remaining for %v\n", remainingTickets, conferenceName)

			firstNames := []string{}
			for _, booking := range bookings {
				var names = strings.Fields(booking)
				firstNames = append(firstNames, names[0])
			}
			fmt.Printf("The first names of bookings are: %v\n ", firstNames)

			noTicketsRemaining := remainingTickets == 0
			if noTicketsRemaining {
				// end program
				fmt.Println("Our conference is booked out")
				break
			}

		} else {
			fmt.Printf("We only have %v ticket remaining, so you can't book %v ticket \n", remainingTickets, userTickets)
		}

	}
}
