package main

import "fmt"

func main() {
	// maps in go lang
	// key value pair
	// key is unique

	//syntax
	lang := make(map[string]string)
	lang["hi"] = "hindi"
	lang["en"] = "english"
	lang["es"] = "spanish"

	fmt.Println("list of lanuages is :-", lang)

	// delete a key value pair
	delete(lang,"es")
	fmt.Println("list of updated one is :-",lang)

	// int and string 
	king:= make(map[int]string)
	king[1] = "hindi"
	king[2] = "english"
	king[3] = "spanish"

	fmt.Println("list of kings is :-",king)
	for key,value := range king{
		fmt.Println("key is",key,"value is",value)
	}

}