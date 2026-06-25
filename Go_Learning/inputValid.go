// package main

// import (
// 	"fmt"
// 	"strings"
// )

// func isinputvalid() {

// 	var userInput string
// 	var userInput2 string
// 	var userlastname string

// 	fmt.Printf("please enter yout anme : \n")
// 	fmt.Scan(&userInput)

// 	fmt.Println("please enter your lasr name: ")
// 	fmt.Scan(&userlastname)

// 	fmt.Print("please enter your email\n")
// 	fmt.Scan(&userInput2)

// 	//input validation
// 	var isValidName = len(userInput) >= 2 && len(userlastname) >= 2
// 	var isValidEmail = strings.Contains(userInput2, "@")

// 	if isValidName && isValidEmail {
// 		fmt.Printf("Your first name is: %v\n", userInput)
// 		fmt.Printf("your last name is %v\n", userlastname)
// 		fmt.Printf("your full name is : %v %v\n", userInput, userlastname)
// 		fmt.Printf("your email is: %v\n", userInput2)
// 	} else {
// 		if !isValidName{
// 			fmt.Println("your firstname and last name shouild be atleast 2 caracter and max 10 caracter")
// 		}
// 		if !isValidEmail {
// 			fmt.Println("your email is invalid")
// 		}
// 		println("your inputs are not valid")
// 	}
// }

// func main4() {//main function is not detactating because os main main function in other files , make it change before runnign\\ng
// 	isinputvalid()
// }
