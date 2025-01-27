package main

import (
	"fmt"
)

/*
Question:
Write a Go program that performs the following tasks using slices:

Create a slice of integers with the values [10, 20, 30, 40, 50, 60, 70, 80, 90, 100].
Prompt the user to enter a starting index and an ending index. Use these values to create a new slice from the existing slice by slicing it.
Validate that the indices are within the range of the original slice.
Calculate and print the sum of all elements in the new slice.
Append the values 200 and 300 to the original slice and print the updated slice.
Print the length and capacity of the original slice before and after the append operation.
*/
func main() {
	var slice = []int{10, 20, 30, 40, 50, 60, 70, 80, 90, 100}
	fmt.Println(slice)

	fmt.Println("Enter s starting and ending index")
	var startIndex int
	var endIndex int
	var newSlice = make([]int, len(slice))
	fmt.Scan(&startIndex, &endIndex)

	if startIndex > len(slice)-1 && startIndex < 0 || endIndex > len(slice)-1 && endIndex < 0 {
		println("Error wrong details")
		return
	} else {
		newSlice = slice[startIndex:endIndex]
		sum := 0
		for i := 0; i < len(newSlice); i++ {
			sum += newSlice[i]
		}
		fmt.Printf("Sum of elements in new slice is %d \n", sum)
		fmt.Printf("len of slice is %d \n", len(slice))
		fmt.Printf("Cap od the slice is %d \n", cap(slice))
		slice = append(slice, 200, 300)
		fmt.Printf("apd:", slice)
		fmt.Printf("len of slice is %d \n", len(slice))
		fmt.Printf("Cap od the slice is %d \n", cap(slice))
	}
}
