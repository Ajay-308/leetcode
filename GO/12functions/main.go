package main

import "fmt"

// a ka type , b ka type , return type
// isko ek or tarah se likh sakte hai

//func add(a,b int) int
func add(a int , b int) int{
	return a+b
}

// another way 
// func lang()(string,string, string){
func lang()(string,string,bool){

	return "hindi", "eng",true
	// agar hume boolean return karna ho kisi ki jagha to change string to bool
}

func main(){
	lang1, lang2,_:= lang()
	fmt.Println(lang1,lang2)

	

}