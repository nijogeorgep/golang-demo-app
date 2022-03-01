package helper

import (
	"fmt"
	"strings"
	"sync"
	"time"
)

func ValidateInputParameters(firstName string, lastName string, email string, userTickets uint, remainingConferenceTickets uint) (bool, bool, bool) {
	isValidName := len(firstName) >= 2 && len(lastName) >= 2
	isValidEmail := strings.Contains(email, "@")
	isValidTicketNumber := userTickets > 0 && userTickets <= remainingConferenceTickets
	return isValidName, isValidEmail, isValidTicketNumber
}

func SendTicket(userTickets uint, firstName string, lastName string, email string, wg sync.WaitGroup) {
	time.Sleep(50 * time.Second)
	var ticket = fmt.Sprintf("%v tickets for %v %v", userTickets, firstName, lastName)
	fmt.Println("#################")
	fmt.Printf("Sending ticket:\n %v \nto email address %v\n", ticket, email)
	fmt.Println("#################")
	wg.Done() // Decrements the waitgroup counter by 1
}
