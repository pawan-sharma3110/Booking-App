package main

import (
	"fmt"
	"strings"
)

func main() {
	var conferenceName string = "Go Conference"
	const conferenceTickets uint = 50
	var remaingTickets uint = 50
	bookings := []string{}
	fmt.Printf("Welcome to our %v booking application \n", conferenceName)
	fmt.Printf("We have total of %v tickets and %v are still available \n", conferenceTickets, remaingTickets)
	fmt.Println("Get your tickets here to attend")
	for {
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

		if userTickets <= remaingTickets {
			remaingTickets = remaingTickets - userTickets
			bookings = append(bookings, firstName+" "+lastName)
			fmt.Printf("The Whole Slice : %v \n", bookings)
			fmt.Printf("The First value : %v  \n", bookings[0])
			fmt.Printf("The Type of Slice :%T \n", bookings)
			fmt.Printf("The lenth of Slice :%v \n", len(bookings))

			fmt.Printf("Thank you %v %v for booking %v tickets. You will recived a confirmation email at %v \n", firstName, lastName, userTickets, email)
			fmt.Printf("%v tickets remaning for %v \n", remaingTickets, conferenceName)
			var shortName = []string{}
			for _, booking := range bookings {
				var name = strings.Fields(booking)
				shortName = append(shortName, name[0])
			}
			fmt.Printf("The first name of booking are: %v\n", shortName)
			if remaingTickets == 0 {
				fmt.Println("Our conference is booked out. Come back next year.")
				break
			} else {
				fmt.Printf(" Dear %v we have only %v tickets reamining , so you cant book %v tickets \n ", firstName, remaingTickets, userTickets)
			}

		}

	}
}
