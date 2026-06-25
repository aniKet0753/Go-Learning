package main

import (
	"fmt"
)
var scanning string
func halper(){
	fmt.Println("this is from hlper function ")
	fmt.Scan(&scanning)
	fmt.Printf("this is the result from halper function and package %v",scanning)

}