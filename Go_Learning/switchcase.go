package main

import (
	"fmt"
)
func main(){
var city  string

	fmt.Println("please enter a city: ")
	fmt.Scan(&city)

	switch city {
	case "kolkata" :
		fmt.Print("welcome to kolkata")
	case "khagaria": 
	fmt.Print("welcome to khagaria")
	case "pune" : 
	fmt.Print("welcome to pune")
	case "banglore" :
		fmt.Print("welcome to banglore") 
	default:
		fmt.Print("this is default print")
	}
	fmt.Println("this is from main funciton")
	halper()
}

