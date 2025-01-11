package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Now()
	fmt.Println(start)
	// lets format the time 

	fmt.Println(start.Format("01-02-2006 monday"))
	// lets get monday or tuesday today 
	fmt.Println(start.Weekday())

	createdDate := time.Date(2021, time.January, 10,23,23,0,0, time.UTC)
	fmt.Println(createdDate.Format("01-02-2006 "))
	fmt.Println(createdDate)
	

}