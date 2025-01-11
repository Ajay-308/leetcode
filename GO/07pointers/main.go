package main

import "fmt"

func main() {

	// pointer is a variable that stores the memory address of another variable
	// & is used to get the memory address of that variable

	// * is used to get the value present at that memory address

	// lets see an example

	var a int = 10
	var ptr = &a

	fmt.Println("value of p is ", ptr)
	fmt.Println("value at p is ", *ptr)

	*ptr = *ptr + 1
	fmt.Println("value of a is ", a)
}