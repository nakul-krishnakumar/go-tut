package main

import (
	"fmt"
	"time"
)

//? WE CAN USE STRUCT EMBEDDINGS TO IMPLEMENT COMPOSITION, INHERITANCE AND OTHER FEATURES OF OOP LANGUAGES

type customer struct {
	name string
	phone uint64
}

type order struct {
	id string
	amount float32
	status string
	createdAt time.Time
	customer // struct embedding
}

func main() {
	newOrder := order{
		id: "2DS",
		amount: 450.44,
		status: "pending",
		customer: customer{
			name: "nakul",
			phone: 8292017483,
		},
	}

	newOrder.customer.name = "amal"

	fmt.Println("New Order:", newOrder)
}