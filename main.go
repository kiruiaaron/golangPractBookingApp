package main

import (
	"fmt"
	"strings"
)

func main() {

	conferenceName := "Go conference"
	const conferenceTickets = 50
	bookings := []string{}
	var remainingTickets uint = 50

	fmt.Printf("conference is %T, remainingTickets is %T, conferenceName is %T\n", conferenceName, conferenceTickets, remainingTickets)

	fmt.Println("Welcome to ", conferenceName, " booking application")
	fmt.Println("we have total of", conferenceTickets, "tickets and ", remainingTickets, "are still available")
	fmt.Printf("We have a total of %v tickets and %v are still available.\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here   ")

	for {

		var firstName string
		var lastName string
		var email string
		var userTickets uint

		//ask user for their name
		fmt.Println("Enter your name")
		fmt.Scan(&firstName)

		fmt.Println("Enter your last name")
		fmt.Scan(&lastName)

		fmt.Println("Enter your email address:")
		fmt.Scan(&email)

		fmt.Println("Enter the number of tickets:")
		fmt.Scan(&userTickets)

		//logical operator

	    isValidName := len(firstName) >= 2 && len(lastName)>=2
        isValidEmail := strings.Contains(email, "@")
		isValidTicketNumber := userTickets > 0 && userTickets <= remainingTickets

       // isValidCity := city == "Singapore" || city == "London"
         
		if isValidName && isValidEmail && isValidTicketNumber   {
			remainingTickets = remainingTickets - userTickets
			//bookings[0] = firstName + " " + lastName
			bookings = append(bookings, firstName+" "+lastName)

			/*fmt.Printf("The whole array: %v\n", bookings)
			fmt.Printf("The first value: %v\n", bookings[0])
			fmt.Printf("slice type: %T \n", bookings)
			fmt.Printf("slice Length: %v\n", len(bookings))
			*/

			fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation email at %v\n", firstName, lastName, userTickets, email)

			fmt.Printf("%v tickets remaining for %v \n", remainingTickets, conferenceName)

			firstNames := []string{}
			for _, booking := range bookings {
				var names = strings.Fields(booking)
				firstNames = append(firstNames, names[0])
			}
			fmt.Printf("The first name of  all bookings: %v\n", firstNames)

			noTicketsRemaining := remainingTickets == 0
			if noTicketsRemaining {
				//end program
				fmt.Println("our conference is booked out. Come back next year.")
				break
			}
		} else  {
			if !isValidName{
				fmt.Println("first name or last name you entered is too short")
			}  
			if !isValidEmail{
				fmt.Println("email address you entered is  invalid")
			}
			if !isValidTicketNumber {
				fmt.Println("number of tickets you entered is invalid")
			}
		}

	}
}
