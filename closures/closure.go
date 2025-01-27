package main

import "fmt"

func main() {
	var funcs []func()

	for i := 0; i < 3; i++ {
		funcs = append(funcs, func() {
			fmt.Println(i)
		})
	}

	for _, f := range funcs {
		f()
	}
}
