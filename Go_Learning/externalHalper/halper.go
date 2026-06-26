package helper

import (
	"fmt"
	"strings"
)

func Helperfunction(firstname string, lastname string , email string)(bool, bool){
	var isvalidname = len(firstname)>=2 && len(lastname)>=2
	var isvalidemail = strings.Contains(email,"@")
	fmt.Printf("this is from helper funtion %v %v %v\n", firstname,lastname,email)
	return isvalidemail , isvalidname
}