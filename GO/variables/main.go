package main

import "fmt"
func main(){
	var name string = "ajay";
	fmt.Println("hello,",name);
	fmt.Printf("hello your type is %T",name);

	// defualt value and same aliases
	var a int 
	fmt.Println("default value of a :",a)
	fmt.Printf("variable of a is type of %T \n",a)


	// implicit type
	// agar hume koi type nhi dete to Go khud hi type decide kar leta hai 
	var b = 10
	fmt.Printf("variable of b is type of %T \n",b)
	// agar ab hum b ki value change karte hai to b ka type bhi change nhi  hoga
	


	///  Go is statically typed and does not allow changing the type of a variable after its declaration.

	// b = "ajay" ---> this will give error because b is int type and we are trying to assign string to it 

	// multiple variable declaration
	var c,d int = 10,20
	fmt.Println("\n",c,d);
}