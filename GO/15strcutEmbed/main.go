package main

import (
	"fmt"
	"time"
)
type parent struct{
	name string
	age int
}
type person struct {
	name      string
	age       int
	gender    string
	createdAt time.Time
	parent
}

func main() {
	// other method to do struct embedding
	parents:= parent{
		name: "ajay",
		age:47,
	}
	p := person{
		name: "vikram",
		age: 25,
		gender: "male",
		//createdAt: time.Time{},
		parent: parents,
		// parent:parent{
		// 	name: "ajay",
		// 	age:47,
		// }
	}
	fmt.Println("parent name:",p.parent.name)
	fmt.Println("name:",p.name)

	fmt.Println(p)

	// in line struct

	lang := struct{
		name  string
		isGood bool
	}{"ajay",true}

	fmt.Println(lang)

}