package main

import (
	"fmt"
	"time"
)

func main() {
	//Write a Go program that takes a number (1 to 7) as input and uses a switch statement to print the corresponding day of the week. If the number is not between 1 and 7, print an error message.
	var input int
	fmt.Println("Enter a number from 1 to 7")
	fmt.Scan(&input)

	switch input {
	case 1:
		fmt.Println(time.Monday)
	case 2:
		fmt.Println(time.Tuesday)
	case 3:
		fmt.Println(time.Wednesday)
	case 4:
		fmt.Println(time.Thursday)
	case 5:
		fmt.Println(time.Friday)
	case 6:
		fmt.Println(time.Saturday)
	case 7:
		fmt.Println(time.Sunday)
	default:
		fmt.Println("Go AWay")
	}

	var num1 int
	var num2 int
	var operator string
	fmt.Println("Enter first Number")
	fmt.Scan(&num1)
	fmt.Println("Enter second Number")
	fmt.Scan(&num2)
	fmt.Println("Enter operator")
	fmt.Scan(&operator)

	switch operator {
	case "+":
		fmt.Printf("Result: %d", num1+num2)

	case "-":
		fmt.Printf("Result: %d", num1-num2)
	case "/":
		fmt.Printf("Result: %d", num1/num2)
	case "*":
		fmt.Printf("Result: %d", num1*num2)

	default:
		fmt.Println("Invalid operator. Please use +, -, *, or /.")
	}
}
