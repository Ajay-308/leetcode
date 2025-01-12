package main

import (
	"fmt"
)

// func sum(nums ...int) int {

// 	total := 0
// 	for _, num := range nums {
// 		total += num

// 	}
// 	return total
// }

func sum(num...interface{})int{
	// interface is a type that defines a set of methods
	// interface --> any type of data 
	total := 0
	// for _,n := range num{
	// 	//total += n
	// }
	return total
}

func main() {

	res := sum(1, 2, 3, 4, 5)
	fmt.Println(res)

}