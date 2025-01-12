package main

import "fmt"

func main() {
	// for loop in go lang , no while

	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}

	// another way of writing for loo
	fmt.Println("another loop")
	for j:= range 11 {
		fmt.Println(j)
	}
}