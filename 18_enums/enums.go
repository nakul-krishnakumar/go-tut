package main

import "fmt"

// -----------------
// enumerated type

type OrderStatus string

const (
	Recieved OrderStatus = "recieved"
	Pending = "pending"
	Delivered = "delivered"
)

//? if we want to use integer enum values:
// type OrderStatus int

// const (
// 	Recieved OrderStatus = iota
// 	Pending 
// 	Delivered
// )
// iota will set Recieved = 0, Pending = 1, Delivered = 2


// -----------------
// order struct

type order struct {
	id int16
	amount float64
	quantity int16
	status OrderStatus
}

//* status setter function
func (o *order) setStatus(status OrderStatus) {
	o.status = status
	fmt.Println("Order status changed to", status)
}

//* order constructor function
func newOrder(id int16, amount float64, quantity int16) *order {
	myOrder := order{
		id: id,
		amount: amount,
		quantity: quantity,
	}

	myOrder.setStatus(Recieved)

	return &myOrder
}

// -----------------

func main() {
	order1 := newOrder(456, 245.42, 45)
	fmt.Println("New Order:", *order1)

	order1.setStatus(Delivered)
	fmt.Println("Updated Order:", *order1)
}