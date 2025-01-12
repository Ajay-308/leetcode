package main

import "fmt"

// slices
// slice are like dynamic arrays with no fixed size
// it provide more functionality than arrays
//
//syntax

func main(){
	//syntax
	var slice [] int = []int {1,2,3,4,5}
	fmt.Println(slice)

	// another way 

	var num = make([]int, 2, 5)
	// here 2 is lenth and 5 is capacity 
	// capacity -> max elements that can be stored in the slice
	num = append(num, 1,2,3,4,5)
	fmt.Println(num)
	fmt.Println(cap(num)) // agar capacity exceed ho jaye to capacity double ho jati hai previous wali ke

	fmt.Println(num[1:3]) // 1 to 3 elements 0 1  , 2 nhi aayega kyuki last wala nhi leta ye


}