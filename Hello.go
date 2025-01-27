package main

import "fmt"

func main() {

	for i := 0; i <= 100; i++ {
		var result = ""
		if i%3 == 0 {
			result += "Fizz"
			fmt.Println(result)
		} else if i%5 == 0 {
			result += "Buzz"
			fmt.Println(result)
		} else if i%3 == 0 && i%5 == 0 {
			result += "FizzBuzz"
			fmt.Println(result)

		} else {
			fmt.Println(i)
		}
	}
	var start int
	var end int
	fmt.Print("Enter start and end : ")
	fmt.Scan(&start, &end)
	sum := 0

	for i := start; i <= end; i++ {
		if i%2 == 0 {
			sum += i
		}

	}
	fmt.Println(sum)

}
