package main

import (
	"fmt"
	"time"
)

/*
	custom class
	interface dependenices inversion
	solid application

	enums: enumerated types
	we implemented enums with help of const

	type MyType string

	type OrderStatus int

	const (
		received orderStatus = iota || "pass same value"
		confirmed
		prepared
		delivered
	)

	
	func changeOrderStatus(status OrderStatus){
		fmt.Println("changing order status to",status)
	}

	func main(){
		changeOrderStatus(received)
	}


	Generics

	func printSlice[T any](item []T){
		for _,i:=range items{
			fmt.Println(i)
		
		}
	}


*/

type paymeter interface {
	pay(amount float32) //return values
	
}

type payment struct{
	gateway paymeter
}

func (p payment) makePayment(amount float32){
	// razorpayment:=razopay{}
	// stripePayment:=stripe{}
	// stripePayment.pay(amount)
	// razorpayment.pay(amount)
	p.gateway.pay(amount)
}

type fakepayment struct {}
func (f fakepayment) pay(amount float32){

	fmt.Println("making payment for testing purpose",amount)
}

type razopay struct{}
func (r razopay) pay(amount float32){

	// logic
	fmt.Println("making payment by razorpay",amount)
}

type stripe struct{}
func (r stripe) pay(amount float32){

	// logic
	fmt.Println("making payment by stripe",amount)
}


type order struct {
	id string
	amount float32
	status string
	createdby time.Time // nanosecond
}

func newOrder(id string,amount float32,status string) *order{
	order:=order{
		id: id,
		amount: amount,
		status: status,
	}

	return &order
}
// receiver type

func (o *order)changeStatus(status string){

	o.status=status
}

// func main(){

// 	myorder:=newOrder("1",34.98,"confirmed")
// 	fmt.Println(myorder)
// 	order := order {
// 		id: "124",
// 		amount: 13454.87,
// 		status: "pending",
// 	}

// 	order.changeStatus("paid")
// 	order.createdby=time.Now()

// 	fmt.Println(order)

// 	// stripePayment:=stripe{}
// 	// razopayPayment:=razopay{}
// 	fakepayments:=fakepayment{}
// 	newpayment :=payment{
// 		// gateway: stripePayment,
// 		// gateway: razopayPayment,
// 		gateway: fakepayments,

// 	}
// 	newpayment.makePayment(100.90)
// }