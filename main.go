package main

import (
	"booking-app/helper"
	"fmt"
	"sync"
	"time"
)

const conferenceTickets uint = 50

var conferenceName string = "Go Conference"
var remaingTickets uint = 50
var bookings = make([]UserData, 0)

type UserData struct {
	firstName       string
	lastName        string
	email           string
	numberOfTickets uint
}

var wg = sync.WaitGroup{}

func main() {

	greetUser()

	firstName, lastName, email, userTickets := getUserInput()
	isValidName, isValidEmail, isValidTicketsNumber := helper.ValidateUserInput(firstName, lastName, email, userTickets, remaingTickets)
	if isValidName && isValidEmail && isValidTicketsNumber {

		bookTickits(userTickets, firstName, lastName, email)
		wg.Add(1)
		go sendTickets(userTickets, firstName, lastName, email)
		frstName := getFirstName()
		fmt.Printf("The first name of booking are: %v\n", frstName)
		if remaingTickets == 0 {
			fmt.Println("Our conference is booked out. Come back next year.")

		}

	} else {
		if !isValidName {
			fmt.Println("Yor entered name is too short")
		}
		if !isValidEmail {
			fmt.Println("Yor entered invalid Email-Id")
		}
		if remaingTickets < userTickets {
			fmt.Printf(" Dear %v we have only %v tickets reamining , so you cant book %v tickets \n ", firstName, remaingTickets, userTickets)
		}
		if !isValidTicketsNumber {
			fmt.Println("Numbert of tickets you entered is invalid ")

		}

	}
	wg.Wait()
}
func greetUser() {
	fmt.Printf("Welcome to %v booking aplication \n", conferenceName)
	fmt.Printf("We have total of %v tickets and %v are still available \n", conferenceTickets, remaingTickets)
	fmt.Println("Get your tickets here to attend")
}
func getFirstName() []string {
	var firstNames = []string{}
	for _, booking := range bookings {

		firstNames = append(firstNames, booking.firstName)
	}
	return firstNames

}

func getUserInput() (string, string, string, uint) {
	var firstName string
	var lastName string
	var email string
	var userTickets uint
	fmt.Println("Enter First Your Name")
	fmt.Scan(&firstName)
	fmt.Println("Enter Last Your Name")
	fmt.Scan(&lastName)
	fmt.Println("Enter Your Email")
	fmt.Scan(&email)
	fmt.Println("Enter how many tickets you wants")
	fmt.Scan(&userTickets)
	return firstName, lastName, email, userTickets
}
func bookTickits(userTickets uint, firstName string, lastName string, email string) {
	remaingTickets = remaingTickets - userTickets

	var userData = UserData{
		firstName:       firstName,
		lastName:        lastName,
		email:           email,
		numberOfTickets: userTickets,
	}

	bookings = append(bookings, userData)
	fmt.Printf("The list of booking is %v \n", userData)
	fmt.Printf("Thank you %v %v for booking %v tickets. You will recived a confirmation email at %v \n", firstName, lastName, userTickets, email)
	fmt.Printf("%v tickets remaning for %v \n", remaingTickets, conferenceName)

}

func sendTickets(userTickets uint, firstname string, lastName string, email string) {
	time.Sleep(10 * time.Second)
	ticket := fmt.Sprintf("%v tickets for %v %v", userTickets, firstname, lastName)
	fmt.Println("#################")
	fmt.Printf("Sending ticket : \n %v to email address %v \n", ticket, email)
	fmt.Println("#################")
	wg.Done()
}
