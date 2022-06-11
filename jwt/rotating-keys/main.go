package main

import (
	"example-jwt/utils"
	"fmt"
)

// TO DO
// create a method to delete the keys that are more than a week old
func main() {
	claims :=utils.UserClaims{
		SessionID: 3,
	}
	fmt.Println("hi there,")
	token, err := utils.CreateToken(&claims)
	fmt.Printf("\nthe token is : %v\n", token)
	if err != nil {
		fmt.Println("some error in creating token", err)
	} else {
		parsedToken, err := utils.ParseToken(token)
		if err != nil {
			fmt.Printf("\nThe error in parsing token : %v\n", err)
		}
		fmt.Println(parsedToken)
	}

}
