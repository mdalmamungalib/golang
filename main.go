package main

import "fmt"

func main() {
	// types of operators
	// 1. Arithmetic Operators --> +, -, *, /, %
	// var num1, num2 int
	// fmt.Printf("num1 = ")
	// fmt.Scan(&num1)

	// fmt.Printf("num2 = ")
	// fmt.Scan(&num2)

	// result := num1 + num2
	// fmt.Printf("%v + %v = %v \n", num1, num2, result)

	// result = num1 - num2
	// fmt.Printf("%v - %v = %v \n", num1, num2, result)

	// result = num1 * num2
	// fmt.Printf("%v * %v = %v \n", num1, num2, result)

	// var result2 = float32(num1) / float32(num2)
	// fmt.Printf("%v / %v = %.2f \n", num1, num2, result2)

	// result = num1 % num2
	// fmt.Printf("%v %% %v = %v \n", num1, num2, result)

	// var base, height, area float32

	// fmt.Printf("Base = ")
	// fmt.Scan(&base)

	// fmt.Printf("Height = ")
	// fmt.Scan(&height)

	// area = 0.5 * base * height
	// fmt.Printf("Area of triangle = %v\n", area)

	var radius, area float32

	fmt.Printf("Enter Radius : ")
	fmt.Scan(&radius)

	area = 3.1416 * radius * radius
	fmt.Printf("Area of circle: %v\n", area)

	// 2. Assignment Operators
	// 3. Unary Operators
	// 4. Relational Operators
	// 5. Logical Operators
	// 6. Bitwise Operators
	// 7. Special Operators
}