package main

import (
	"fmt"
	"time"
)

/*
	custom class
	

	interface
*/

type payment struct{}

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

func main(){

	myorder:=newOrder("1",34.98,"confirmed")
	fmt.Println(myorder)
	order := order {
		id: "124",
		amount: 13454.87,
		status: "pending",
	}

	order.changeStatus("paid")
	order.createdby=time.Now()

	fmt.Println(order)
}