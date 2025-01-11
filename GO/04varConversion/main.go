package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main(){

	welcome := "welcome to user input"
	fmt.Println(welcome)

	// var input as string

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("enter rating for ajay 1-10:")

	// now we have to convert string to float64 or int
	// we can use strconv pakcage for this

	input,_:= reader.ReadString('\n')
	rating,err := strconv.ParseFloat(strings.TrimSpace(input),64)
	// if we want to convert strint to int we can use strconv.Atoi
	// hume yaha trim karne ki jarurt isliye padi kyuki user ne enter press kiya tha to new line character bhi aa gya tha 
	// jisse parsing to float mai error aa raha tha
	if err != nil{
		fmt.Println(err)

	}else{
		
		fmt.Println("your rating is :",rating + 1)
	}
	
}