package main

import "fmt"

// interface --> ka interface kisi bhi type ki value ko hold kar sakta hai
// iska use hum polimorphism ke liye karte hai
// jaise ki agar hume ek function likhna hai jo kisi bhi type ki value ko accept kare to hum uske liye interface ka use kar sakte hai

// open close principle
// open for extension but close for modification --> current code doesn't follow this rule until we use interface
// let use it

type paymenter interface{ // er in end  is a convention to write interface name 
	pay(amount int) // ye same hone chiye sare gatewaye ka jo bhi params hai 
}
type payment struct{
	// ab hume payment ke liye kisi bhi type ke gateway ka use karna hai to hume ek interface banani pagegi jisme ek method hoga pay
	// jo ki bhut sare gateway ke liye bhut sare banane padenge jaise razorpay,stripe,fakepay etc
	// isse acha hai ki hum ek interface bana le jisme ek method pay hoga jo ki kisi bhi gatewaye ke liye common hoga 
	

	gateway paymenter

}

func(p payment) pay(amount int){
	// logic
	p.gateway.pay(amount)
	fmt.Println("payment done of amount:",amount)
}

type razorpay struct{}

func(r razorpay)pay(amount int){
	//logic
	fmt.Println("payment done by razorpay of amount:",amount)
}

type stripe struct{}

func(r stripe)pay(amount int){
	//logic
	fmt.Println("payment done by stripe of amount:",amount)
}


type fakepay struct{}

func(r fakepay)pay(amount int){
	//logic
	fmt.Println("payment done by fakepay of amount:",amount)
}



func main(){
	newPayment := payment{
		gateway:fakepay{},
	}
	newPayment.pay(100)

}