package main

import (
	"fmt"
)

type name struct {
	message   string // poneNumber int
	sender    user
	recipient user
}
type user struct {
	name        string
	phoneNumber int
}

// function signature
func sub(x int, y int) int {
	return x - y
}
func add(x, y int) int {
	return x + y
}

func main() {
	// this project is the basic syntax of using golang
	// if you are note using a varaible in go land you indcate it with an underscore _
	// var smsSendingLimit int
	// var costPerSMS float64
	// var hasPermission bool
	// var username string
	// fmt.Printf(
	// 	"%v %f %v %q \n",
	// 	smsSendingLimit,
	// 	costPerSMS,
	// 	hasPermission,
	// 	username,
	// )
	// // var username string = "taiye"
	// // var password int = 44
	// fmt.Println("hello kehinde", username)
	// congrats := "Happy Birthday"
	// fmt.Println(congrats)

	// penniesPerText := 2.0
	// fmt.Println("This type of is of", penniesPerText)
	//    AccountAge:= 2.6
	// accountAgeInt:= int(AccountAge)
	// fmt.Println("The age is of",accountAgeInt)
	// mileage, company := .23, "10x Developers"
	// fmt.Println(mileage, company)

	const AFullSemseter = "full"
	const premeiumSemeter = "half"
	fmt.Println("this is a basic plan", premeiumSemeter)
	fmt.Println("this is a premium plan", AFullSemseter)

	// conditional in Golang
	messageLen := 10
	maxMessageLen := 20
	fmt.Println("Trying to print the length", messageLen, maxMessageLen)
	if messageLen <= 5 {
		fmt.Println("The message is too long")
	}

}
