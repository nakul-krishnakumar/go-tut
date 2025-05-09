package main

import "fmt"

//! WHEN TO USE INTERFACES?
/*//* Use interfaces when:
		- You want flexibility (plug-and-play behavior).

		- You want to follow dependency inversion — depend on abstractions, not implementations.

		- You want to mock behavior (e.g., in tests).
*/


//? NAMING CONVENTION -> end with 'er'
type paymenter interface {
	pay(amount float32) bool
}
//* Here, if we did not use an interface, we will have to explicitly hard code which payment gateway(stripe or razorpay) at initialization itself.
//? When we use interface, we can setup a payment Interface (paymenter <IPaymentGateway>) which let us choose which payment method to choose
//! For this payment interface to work, all the payment gateways (stripe, razorpay) should have the `pay(amount float32) bool` function

type payment struct {
	gateway paymenter
}

func (p payment) makePayment(amount float32) {
	p.gateway.pay(amount)
}

// ----------

type stripe struct {}

func (s stripe) pay(amount float32) bool {
	fmt.Println(amount, "paid through stripe!")
	return true
}

// ----------

type razorpay struct {}

func (r razorpay) pay(amount float32) bool {
	fmt.Println(amount, "paid through razorpay!")
	return true
}

// ----------

func main() {
	razorGW := razorpay{}
	newPayment := payment{gateway: razorGW}
	newPayment.makePayment(5454.33)
}