package main

import (
	"bufio"
	"fmt"
	"os"
)

func main(){
	welcome := "welcome to user input "
	fmt.Println(welcome)
	
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("enter your name:")

	// in go lang we don't have try catch block 
	// we have comma ok in go 

	//agar sirf input lena hai to input, _ 
	// agar error bhi check karna hai to input,err


	input,_ := reader.ReadString('\n')
	fmt.Println("hello",input)
	fmt.Printf("your type is %T",input)


}