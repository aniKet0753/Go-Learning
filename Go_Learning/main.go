package main

import "fmt"

func main() {
	var conferenceName = "conference booking application"
	const conferenceTicket = 50
	var remaningTickets uint = 50
	fmt.Printf("Welcome to Our %v booking application\n", conferenceName)
	fmt.Printf("we have total of  %v tickets and %v are still avaliable\n", conferenceTicket, remaningTickets)
	fmt.Println("Get Your Ticket here to attend")

	var firstName string
	var lastName string
	var eamil string
	var userTickets uint
	//ask username
	fmt.Println("Enter your first name: ")
	fmt.Scan(&firstName)//& defines pointer , the adddress of the memory where the userName value is stored.

	fmt.Println("Enter your lastName")
	fmt.Scan(&lastName)

	fmt.Println("Enter your email:")
	fmt.Scan(&eamil)

	fmt.Println("enter number of tickets: ")
	fmt.Scan(&userTickets)

	remaningTickets = remaningTickets - userTickets//type mismach between both rmainingTicket and remaingTicket


	fmt.Printf("Tkank you %v %v for booling %v ticket, you will recive a conformation email at %v\n",firstName, lastName, userTickets, eamil)
	fmt.Printf(" %v is remaning tickets for the confrenece Tickets %v \n",remaningTickets,conferenceTicket)
}
