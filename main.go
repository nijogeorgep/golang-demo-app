package main

import (
	"booking-app/booking"
	"booking-app/helper"
	"fmt"
	"sync"
)

//Package Level Variables - Accessible to All
const conferenceTickets int = 50

//Type Inference in Variabls, Datatypes
var conferenceName = "Go Conference"

var remainingConferenceTickets uint = 50

//Go goRoutine WaitGroup for Threads over Green Threads
var wg = sync.WaitGroup{}

func main() {

	//fmt.Printf("conferenceName is %T,conferenceTickets is %T,remainingConferenceTickets is %T\n", conferenceName, conferenceTickets, remainingConferenceTickets)
	greetUsers()

	firstName, lastName, email, userTickets := getUserInputs()
	isValidName, isValidEmail, isValidTicketNumber := helper.ValidateInputParameters(firstName, lastName, email, userTickets, remainingConferenceTickets)

	if isValidName && isValidEmail && isValidTicketNumber {
		//Book Tickets
		booking.BookTickets(userTickets, firstName, lastName, email, conferenceName, remainingConferenceTickets)

		wg.Add(1)                                                         // sets the number of goroutines
		go helper.SendTicket(userTickets, firstName, lastName, email, wg) //Go Coroutine
		//Print FirstNames
		firstNames := booking.GetFirstnames()
		fmt.Printf("All current bookings %v\n", firstNames)

		if remainingConferenceTickets == 0 {
			//end program
			fmt.Printf("%v is booked out, come back next year\n", conferenceName)
			//break
		}
	} else {
		if !isValidName {
			fmt.Println("FirstName and LastName entered is too short")
		}
		if !isValidEmail {
			fmt.Println("Your email is not in a valid format")
		}
		if !isValidTicketNumber {
			fmt.Println("Number tickets entered is Invalid")
		}
		fmt.Println("Your input data is invalid! Please try Again.")
		//fmt.Printf("We only have %v tickets remaining, so you can't book %v tickets\n", remainingConferenceTickets, userTickets)
		//continue
	}
	wg.Wait() // blocks until the WaitGroup counter is 0
}

func greetUsers() {
	//Pointers
	fmt.Println(&conferenceName)
	//https://pkg.go.dev/fmt@go1.17.6
	fmt.Printf("Welcome to %v Booking Application\n", conferenceName)
	fmt.Printf("We have total of %v tickets and %v are still available!\n", conferenceTickets, remainingConferenceTickets)
	fmt.Println("Get your tickets here to attend!")
}

func getUserInputs() (string, string, string, uint) {
	var firstName string
	var lastName string
	var email string
	var userTickets uint
	//Ask user inputs
	//Pointers
	fmt.Println("Enter your firstName: ")
	fmt.Scan(&firstName)

	fmt.Println("Enter your lastName: ")
	fmt.Scan(&lastName)

	fmt.Println("Enter your email address: ")
	fmt.Scan(&email)

	fmt.Println("Enter the number of tickets: ")
	fmt.Scan(&userTickets)

	return firstName, lastName, email, userTickets
}
