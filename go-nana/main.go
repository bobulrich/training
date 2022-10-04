package main

import (
	"fmt"
	"strings"
)

func main() {
	conferenceName := "Go Conference"
	const conferenceTickets int = 50
	var remainingTickets int = 50
	var bookings = []string{}

	fmt.Printf("conferenceTickets is %T, remainingTickets is %T, conferenceName is %T\n", conferenceTickets, remainingTickets, conferenceName)
	fmt.Printf("Welcome to %v booking application\n", conferenceName)
	fmt.Printf("We have total of %v tickets and %v remaining tickets\n", conferenceTickets, remainingTickets)
	fmt.Printf("Get your tickets here to attend\n")

	for {
		var firstName string
		var lastName string
		var email string
		var userTickets int

		fmt.Printf("Enter your first name: \n")
		fmt.Scan(&firstName)

		fmt.Printf("Enter your last name: \n")
		fmt.Scan(&lastName)

		fmt.Printf("Enter your email: \n")
		fmt.Scan(&email)

		fmt.Printf("Enter number of tickets desired: \n")
		fmt.Scan(&userTickets)

		isValidName := len(firstName) >= 2 && len(lastName) >= 2
		isValidEmail := strings.Contains(email, "@")
		isValidUserTickets := userTickets > 0 && userTickets <= remainingTickets
		// isValidCity := strings.ToLower(city) == "singapore" || strings.ToLower(city) == "london"

		if isValidEmail && isValidUserTickets && isValidName {
			remainingTickets = remainingTickets - userTickets

			bookings = append(bookings, firstName+" "+lastName)

			fmt.Printf("These are all our bookings: %v\n", bookings)
			fmt.Printf("User %v %v booked %v tickets.\n", firstName, lastName, userTickets)
			fmt.Printf("Thank you for purchasing %v tickets.\nYou will receive confirmation email at: %v\n", userTickets, email)
			fmt.Printf("Remaining tickets: %v\n", remainingTickets)

			firstNames := []string{}
			for _, booking := range bookings {
				names := strings.Fields(booking)
				firstName = names[0]
				firstNames = append(firstNames, firstName)
			}

			fmt.Printf("The first names of bookings are: %v\n", firstNames)

			noTicketsRemaining := remainingTickets <= 0
			if noTicketsRemaining {
				// end program
				fmt.Printf("Our conference is booked out. Come back next year!\n")
				break
			}
		} else {
			if !isValidName {
				fmt.Println("First name or last name you entered is too short")
			}
			if !isValidEmail {
				fmt.Println("Email format is invalid")
			}
			if !isValidUserTickets {
				fmt.Println("Number of tickets is invalid, must be >0")
			}
			// fmt.Println("Your input data is invalid, try again...")
		}
	}
}
