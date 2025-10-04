package main

import "fmt"

func main() {
	age := 20 // Declare and initialize age variable
	if age > 18 {
		// Age is greater than 18
		fmt.Println("You are eligible to get married.")
	} else if age < 18 {
		// Age is less than 18
		fmt.Println("Sorry! you are not eligible to get married")
	} else {
		// Age is exactly 18
		fmt.Println("Oh boy! You are a teenager.")
	}
}
