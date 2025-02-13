<<<<<<< HEAD
package main

import (
	"fmt"
)

// standard module to primt things the console

func main() {
	// var timname string = "Kehinde"
	// //tim_name2 = "tola"
	// //var number uint16 = 260
	// number := 99
	//   fmt.Println()

	// fmt.Println("Hello world", number, timname)
	// fmt.Printf("%T")
	//var x string = fmt.Sprintf("")
	//var scanner := bufio.NewScanner();
	//var scanner bufio.Scanner

	// FOR INPUT VALUES
	// scanner := bufio.NewScanner(os.Stdin)
	// fmt.Printf("Type the year you were born")
	// scanner.Scan()
	// input, _ := strconv.ParseInt(scanner.Text(), 10, 64)
	// fmt.Printf("You will be %d years old at the end of 2020", 2020-input)
	// VARIABLES
	// math.NaN()
	// var num1 int = 8
	// var num2 int = 4
	// answer := num1 / num2
	// fmt.Printf("%d", answer)

	// Conditional Expressions boolean
	// x := 5
	// y := 6
	// val := x < 5
	// fmt.Printf("%t", val)
	// If else statements
	fmt.Println("Hello world")
	//name := "tim"
	age := 20

	if age >= 18 {
		fmt.Println("You can hear the ride")
	} else if age >= 14 {
		fmt.Println("you can ride with your parents")
	} else {
		fmt.Println("you can not ride!")
	}

	if age >= 18 {
		fmt.Println("you can ride a car")
	} else {
		fmt.Println("After it")
		fmt.Printf("Wait %d years", 18-age)
	}
}
=======
package main t

import "fmt" // standard module to print things the console

func main() {
	var tim_name2 string = "Kehinde"
	tim_name2 = "tola"
	var number uint8 = 260

	fmt.Println("Hello world", number)
}
>>>>>>> d8adf5213a14de53f718677d0d0a3c45b6765184
