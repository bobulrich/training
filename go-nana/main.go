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

	greetUsers(conferenceName, conferenceTickets, uint(remainingTickets))

	fmt.Printf("conferenceTickets is %T, remainingTickets is %T, conferenceName is %T\n", conferenceTickets, remainingTickets, conferenceName)
	fmt.Printf("Welcome to %v booking application\n", conferenceName)
	fmt.Printf("We have total of %v tickets and %v remaining tickets\n", conferenceTickets, remainingTickets)
	fmt.Printf("Get your tickets here to attend\n")

	for {
		firstName, lastName, email, userTickets := getUserInput()
		isValidName, isValidEmail, isValidTicketNumber := validateUserInput(firstName, lastName, email, userTickets, uint(remainingTickets))
		
		if isValidEmail && isValidUserTickets && isValidName {
			remainingTickets = remainingTickets - userTickets
			
			bookings = append(bookings, firstName+" "+lastName)
			
			fmt.Printf("These are all our bookings: %v\n", bookings)
			fmt.Printf("User %v %v booked %v tickets.\n", firstName, lastName, userTickets)
			fmt.Printf("Thank you for purchasing %v tickets.\nYou will receive confirmation email at: %v\n", userTickets, email)
			fmt.Printf("Remaining tickets: %v\n", remainingTickets)
			
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
			
			firstNames := getFirstNames(bookings)
			fmt.Println(firstNames)
			
			greetUsers(conferenceName, conferenceTickets, remainingTickets)

			city := "london"
			
			switch city {
			case "london":
				fmt.Print("Execute london code")
			case "singapore":
				fmt.Print("Execute singapore code")
			case "berlin":
				fmt.Print("Execute berlin code")
			case "mexico city":
				fmt.Print("Execute mexico city code")
			case "hong kong":
				fmt.Print("Execute hong kong code")
			default:
				fmt.Print("No valid city selected")
				
			}
		}
		
		func greetUsers(confName string, confTickets int, remainingTickets uint) {
			fmt.Printf("Welcome to %v booking application\n", confName)
			fmt.Printf("We have a total of %v tickets and %v are still available.\n", confTickets, remainingTickets)
		}
		
		func getFirstNames(bookings []string) []string {
			firstNames := []string{}
			for _, booking := range bookings {
				names := strings.Fields(booking)
				firstNames = append(firstNames, names[0])
			}
			return firstNames
	}
}
		
func validateUserInput(firstName string, lastName string, email string, userTickets int, remainingTickets uint) (bool, bool, bool) {
	isValidName := len(firstName) >= 2 && len(lastName) >= 2
	isValidEmail := strings.Contains(email, "@")
	isValidUserTickets := userTickets > 0 && userTickets <= int(remainingTickets)
	return isValidName, isValidEmail, isValidUserTickets
}

func getUserInput() (string, string, string, int) {

	fmt.Printf("Enter your first name: \n")
	fmt.Scan(&firstName)

	fmt.Printf("Enter your last name: \n")
	fmt.Scan(&lastName)

	fmt.Printf("Enter your email: \n")
	fmt.Scan(&email)

	fmt.Printf("Enter number of tickets desired: \n")
	fmt.Scan(&userTickets)

	return &firstName, &lastName, &email, &userTickets
}