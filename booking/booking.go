package booking

import (
	"fmt"
)

var GlobalVariable = "Global Variable"

//Slice of Empty UserData Array
var ticketBookings = make([]UserData, 0)

//Alternative for Classes in Golang
type UserData struct {
	firstName       string
	lastName        string
	email           string
	numberOfTickets uint
}

func BookTickets(userTickets uint, firstName string, lastName string, email string, confName string, remainingConferenceTickets uint) {
	remainingConferenceTickets = remainingConferenceTickets - userTickets

	//Create a map for a user
	//var userData = make(map[string]string) //creates an empty map
	var userData = UserData{
		firstName:       firstName,
		lastName:        lastName,
		email:           email,
		numberOfTickets: userTickets,
	}

	ticketBookings = append(ticketBookings, userData)
	fmt.Printf("List of Bookings %v\n", ticketBookings)

	fmt.Printf("Thank You %v %v for booking %v tickets. You will recieve a confirmation email at %v\n", firstName, lastName, userTickets, email)
	fmt.Printf("%v tickets remaining for %v\n", remainingConferenceTickets, confName)
	//return TicketBookings
}

func GetFirstnames() []string {
	//slice - growing array
	firstNames := []string{}
	//for index, booking := range ticketBookings {
	for _, booking := range ticketBookings {
		//var names = strings.Fields(booking) //used for spliting the name based on empty space
		firstNames = append(firstNames, booking.firstName)
	}
	//return firstnames
	return firstNames
}
