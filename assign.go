package main

import "fmt" // to be able to print syntax 

func main() {
	sendSoFar := 430
	const sendToAdd = 25
	sendSoFar = incrementSend(sendSoFar, sendToAdd)
	fmt.Printf("Total sends so far: %d\n", sendSoFar)

	fmt.Println("Hello, World!")

	// func getNames() (string string)  {
	// 	return "Kehinde", "Doe"
	// }

}

func incrementSend(sendSoFar, sendToAdd int) int {
	sendSoFar = sendSoFar + sendToAdd
	return sendSoFar
}
