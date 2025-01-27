//Write a Go program that takes a number as input and prints its multiplication table up to 10.

package main

import "fmt"

func main() {
	var input int
	fmt.Println("Please enter a number to generate the times table")
	fmt.Scan(&input)
	for i := 1; i < 11; i++ {
		fmt.Println(fmt.Sprintf(" %d x %d = %d", input, i, input*i))
	}

}
