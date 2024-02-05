package main

import (
	"fmt"
	"time"
	//"strconv"
)

// Package Level variables
// Variables define outside the all the function
const conferenceTickets = 50

var conferenceName = "Go conference"
//var bookings = make([]map[string]string, 0)
var bookings = make([]UserData, 0)
var remainingTickets uint = 50



type UserData struct {
	firstName string
	lastName string
	email string
	numberOfTickets uint
}





func main() {

	greetUsers()

	//fmt.Printf("conference is %T, remainingTickets is %T, conferenceName is %T\n", conferenceName, conferenceTickets, remainingTickets)

	for {

		firstName, lastName, email, userTickets := getUserInputs()

		//logical operator

		isValidName, isValidEmail, isValidTicketNumber :=validateUserInput(firstName, lastName, email, userTickets)

		// isValidCity := city == "Singapore" || city == "London"

		if isValidName && isValidEmail && isValidTicketNumber {

			bookTicket(userTickets, firstName, lastName, email)
			go sendTicket(userTickets, firstName, lastName, email)
 
			//call function print firstNames
			firstNames := getFirstNames()
			fmt.Printf("The first names of the booking are: %v\n", firstNames)

			noTicketsRemaining := remainingTickets == 0
			if noTicketsRemaining {
				//end program
				fmt.Println("our conference is booked out. Come back next year.")
				break
			}
		} else {
			if !isValidName {
				fmt.Println("first name or last name you entered is too short")
			}
			if !isValidEmail {
				fmt.Println("email address you entered is  invalid")
			}
			if !isValidTicketNumber {
				fmt.Println("number of tickets you entered is invalid")
			}
		}

	}

	//switch statement just practice
	/*city := "London"

		switch city {
		case "New York":
			//execute  code for booking new york conference tickets

		case "Singapore","Hong Kong":
		   //execute code for booking Singapore  and Hong Kong conference tickets

		case "Berlin","London":
			//some code here for Berlin and London

		case "Mexico City ":
			//some code

	    default:
			fmt.Println("no Valid city selected")


		}*/

}

//CREATING OTHER FUNCTIONS

func greetUsers() {
	fmt.Printf("Welcome to our %v booking application\n ", conferenceName)
	fmt.Printf("We have a total of %v tickets and %v are still available.\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here  to attend ")
}

func getFirstNames() []string {
	firstNames := []string{} 
	for _, booking := range bookings {
		firstNames = append(firstNames, booking.firstName)
	}
	return firstNames

}

func getUserInputs() (string, string, string, uint) {
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

	return firstName, lastName, email, userTickets
}

func bookTicket(userTickets uint, firstName string, lastName string, email string) {
	remainingTickets = remainingTickets - userTickets


	//create a map for user
   // var userData =make( map[string]string)
    var userData =UserData {
		firstName: firstName,
		lastName: lastName,
		email:email,
		numberOfTickets: userTickets,
	}

	/*
	userData["firstName"] = firstName
	userData["lastName"]= lastName
    userData["email"]= email
	userData["numberOfTickets"]=strconv.FormatUint(uint64(userTickets), 10)
   */

	//bookings[0] = firstName + " " + lastName
	bookings = append(bookings, userData)

	fmt.Printf("List of bookings is %v\n", bookings)

	/*fmt.Printf("The whole array: %v\n", bookings)
	fmt.Printf("The first value: %v\n", bookings[0])
	fmt.Printf("slice type: %T \n", bookings)
	fmt.Printf("slice Length: %v\n", len(bookings))
	*/

	fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation email at %v\n", firstName, lastName, userTickets, email)

	fmt.Printf("%v tickets remaining for %v \n", remainingTickets, conferenceName)

}


func sendTicket(userTickets uint, firstName string, lastName string, email string){

	time.Sleep(10 * time.Second)

	var ticket = fmt.Sprintf("%v tickets for %v %v ", userTickets, firstName, lastName)
	fmt.Println("#########")
	fmt.Printf("Sending ticket:\n %v \n to email address %v\n", ticket, email)
}