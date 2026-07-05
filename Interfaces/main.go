package main

import "fmt"

/*

Interfaces: Imp for writing scalable, maintainable reusable code.
- Real world example -> payment gateway.
- Problem with this => new Requirement come, we need to add stripe paymentGateWay
- Interfaces: It is a contract



*/

type PaymentInterface interface {
	// all methods inside interface are compulsory to implement
	pay(amount float32)
	// refund(account string, amount float32)
}




type Payment struct {
	gateway PaymentInterface
}

// changing main code -> violet open close principle -> methods are open for extension, but close for modification
func (p Payment) makePayment(amount float32) {
	// create payment
	// razorpayGt := Razorpay{}
	// razorpayGt.pay(amount)

	// stripePaymentGT := Stripe{}
	// stripePaymentGT.pay(amount)

	p.gateway.pay(amount)
}




// Razorpay integration
type Razorpay struct{}

func (r Razorpay) pay(amount float32) {
	// logic to make payment
	fmt.Println("Using razorpay, Payment successfully Done!🎉", amount)
}



// Stripe integration
type Stripe struct{}

func (r Stripe) pay(amount float32) {
	// logic to make payment
	fmt.Println("Using stripe, Payment successfully Done!🎉", amount)
}




// future need: Added paypal paymentGateway
type Paypal struct{}

func (p Paypal) pay(amount float32) {
	fmt.Println("Using Paypal, Payment successfully Done!🎉", amount)
}



func main() {

	// newPayment := Payment{}
	// newPayment.makePayment(100.000)

	// razorpayPaymentGw := Razorpay{}
	// stripePaymentGw := Stripe{}
	paypalPaymentGw := Paypal{}

	newPayment := Payment{
		gateway: paypalPaymentGw,
	}

	newPayment.makePayment(2000)

}
