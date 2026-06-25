// package main

// import (
// 	"fmt"
// 	"strings"
// )

// func main2() {
// 	var conferenceName = "conference booking application"
// 	const conferenceTicket = 50
// 	var remaningTickets uint = 50
// 	fmt.Printf("Welcome to Our %v booking application\n", conferenceName)
// 	fmt.Printf("we have total of  %v tickets and %v are still avaliable\n", conferenceTicket, remaningTickets)
// 	fmt.Println("Get Your Ticket here to attend")

// 	var bookings []string

// 	var firstName string
// 	var lastName string
// 	var eamil string
// 	var userTickets uint
// 	for {
// 		//ask username
// 		fmt.Println("Enter your first name: ")
// 		fmt.Scan(&firstName) //& defines pointer , the adddress of the memory where the userName value is stored.

// 		fmt.Println("Enter your lastName")
// 		fmt.Scan(&lastName)

// 		fmt.Println("Enter your email:")
// 		fmt.Scan(&eamil)

// 		fmt.Println("enter number of tickets: ")
// 		fmt.Scan(&userTickets)
		
// 		if userTickets <= remaningTickets {

// 			remaningTickets = remaningTickets - userTickets //type mismach between both rmainingTicket and remaingTicket

// 			// bookings[0] = firstName +" "+ lastName
// 			bookings = append(bookings, firstName+" "+lastName)

// 			fmt.Printf(" the whole bookingslice/: %v\n", bookings)
// 			fmt.Printf(" the booking person first name %v\n", bookings[0])
// 			// fmt.Printf(" length of the array  %v\n", len(bookings))
// 			fmt.Printf("these ae all your bookings %v\n", bookings)

// 			SliceFirstname := []string{}       //created slice
// 			for _, element := range bookings { //slice index and element insid it
// 				var elementname = strings.Fields(element) //libreary for making ["ankit ","kumaar"] to ankit kumar
// 				SliceFirstname = append(SliceFirstname, elementname[0])
// 			}
// 			fmt.Printf("firs name of boking %v\n", SliceFirstname)

// 			notticketRemaning := remaningTickets == 0
// 			if notticketRemaning {
// 				//end program
// 				fmt.Printf("confrenece is booked out , comeback next year ")
// 				break
// 			}

// 			fmt.Printf("Tkank you %v %v for booling %v ticket, you will recive a conformation email at %v\n", firstName, lastName, userTickets, eamil)
// 			fmt.Printf(" %v is remaning tickets for the confrenece Tickets %v \n", remaningTickets, conferenceTicket)

// 		} else {
// 			fmt.Printf("%vThis is not valid, we only hve %v please book within the avlable ticket avlabelity\n", userTickets, remaningTickets)

// 		}

// 	}

// }
