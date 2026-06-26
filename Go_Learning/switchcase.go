package main

import (
	"Go-Learning/externalHalper"
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
	

var firstname string
var lastname string 
var email string


	fmt.Println("inter your first name ")
	fmt.Scan(&firstname)
	fmt.Println("inter your first name ")
	fmt.Scan(&lastname)
	fmt.Println("inter your email ")
	fmt.Scan(&email)
	isvalidemail , isvalidname := helper.Helperfunction(firstname,lastname,email) 
	if isvalidemail && isvalidname {
		fmt.Printf("this input is valid from switch package %v %v\n",isvalidemail,isvalidname)
	}

}

