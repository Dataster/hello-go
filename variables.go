package main

import "fmt"

func main() {
	// Declare a variable
	var name string
	// Assign a value to the variable
	name = "John Doe"
	// Print the variable
	fmt.Println("Hello, ", name)
	// Declare and assign a variable in one line
	var age int = 30
	// Print the variable
	fmt.Println("Age: ", age)
	// Declare and assign a variable with type inference
	country := "USA"
	// Print the variable
	fmt.Println("Country: ", country)
}
