package main

import "fmt"

func main(){

	fmt.Println("Arrays in Go Lang")
	//types of arrays

	fruits := [4]string{"apples", "oranges", "bananas"}
	fmt.Println(fruits)

	// anoter way

	var cars [3] string
	cars[0]="audi"
	cars[2]= "bmw"
	fmt.Println(cars)
}