package main

import "fmt"

func main() {

	var fullName string
	var num1, num2 int
	// var age int
	// var gpa float32
	fmt.Print("Enter your name: ")
	fmt.Scan(&fullName)
	fmt.Print("Please inter your 2 number: ")
	fmt.Scan(&num1, &num2)


	fmt.Printf("%v is a student\n", fullName)
	fmt.Printf("num1 = %v, num2 = %v \n", num1, num2)
}