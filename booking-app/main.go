package main

import "fmt"

func main() {
	var conferenceName = "Go Conference"
	const conferenceTickets = 50
	var remaningTickets = 50

	fmt.Println("Welcome to our", conferenceName, "booking application")
	fmt.Println("We have total of", conferenceTickets, "tickets and", remaningTickets, "are still available.")
	fmt.Println("Get your tickets here to attend")

	var userName string
	var userTickets int

	userName = "Makarand"
	userTickets = 2
	fmt.Println(userName)
	fmt.Println(userTickets)

}
