package main

import "fmt"

func main() {
	var conferenceName string = "Go Conference"
	const conferenceTickets uint = 50
	var remaingTickets uint = 50
	var bookings [50]string
	fmt.Printf("Welcome to our %v booking application \n", conferenceName)
	fmt.Printf("We have total of %v tickets and %v are still available \n", conferenceTickets, remaingTickets)
	fmt.Println("Get your tickets here to attend")

	var firstName string
	var lastName string
	var email string
	var userTickets uint
	fmt.Println("Enter First Your Name")
	fmt.Scan(&firstName)
	fmt.Println("Enter Last Your Name")
	fmt.Scan(&lastName)
	fmt.Println("Enter how many tickets you wants")
	fmt.Scan(&userTickets)
	fmt.Println("Enter Your Email")
	fmt.Scan(&email)
	remaingTickets = remaingTickets - userTickets
	bookings[0] = firstName + " " + lastName
	fmt.Printf("The Whole Array : %v \n", bookings)
	fmt.Printf("The First value : %v  \n", bookings[0])
	fmt.Printf("The Type of Array :%T \n", bookings)
	fmt.Printf("The lenth of Array :%v \n", len(bookings))

	fmt.Printf("Thank you %v %v for booking %v tickets. You will recived a confirmation email at %v \n", firstName, lastName, userTickets, email)
	fmt.Printf("%v tickets remaning for %v \n", remaingTickets, conferenceName)
}
