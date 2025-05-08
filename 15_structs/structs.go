package main

import (
	"fmt"
	"time"
)

//? STRUCTS are used to make custom data types
type order struct {
	id        string
	amount    float32
	status    string
	createdAt time.Time // nanosecond precision
}

//* CONSTRUCTOR FOR STRUCT
//? convention is to name it as new<struct_name>, eg: 'newOrder' or 'NewOrder' 
func newOrder(id string, amount float32, status string) *order {
	myOrder := order{
		id: id,
		amount: amount,
		status: status,
	}

	return &myOrder
}

//* RECEIVER TYPE
// func (o order) changeStatus(status string) { -> Here, a copy of order struct is passed
func (o *order) changeStatus(status string) {
	//* convention is to set object name as first letter of the struct, here 'o'
	o.status = status
}

//? Here passing order struct as both pointer or normal will work as we are just getting the attribute and not updating, so there is no need of passing by reference (pointer)
// func (o *order) getAmount() float32 { 
func (o order) getAmount() float32 { 
	return o.amount
}

func main() {
	//? if you do not set any field, it will be zero value by default
	//? int => 0, float => 0, string => "", bool => false
	myOrder := order{
		id: "1AB",
		amount: 250.00,
		status: "received",
	}

	myOrder.createdAt = time.Now()
	fmt.Println("Struct Order:", myOrder)
	myOrder.changeStatus("pending")
	fmt.Println("Struct Order:", myOrder)

	fmt.Println("Amount:", myOrder.getAmount())

	// using constructor function
	myOrder2 := *newOrder("2DC", 231.22, "recieved")
	fmt.Println("New Order:", myOrder2)

	//* UNNAMED STRUCTS
	lang := struct {
		name string
		isGood bool
	} {"Golang", true}

	fmt.Println("Unnamed struct:", lang)
}
