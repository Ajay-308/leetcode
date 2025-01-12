package main

import (
	"fmt"
	"time"
)

type person struct {
	name      string
	age       int
	gender    string
	createdAt time.Time
}

func main() {
	p := person{
		name: "ajay",
		age: 25,
		gender: "male",
		//createdAt: time.Time{},

	}
	fmt.Println("name:",p.name)

	fmt.Println(p)

	// in line struct

	lang := struct{
		name  string
		isGood bool
	}{"ajay",true}

	fmt.Println(lang)

}