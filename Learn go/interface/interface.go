//TODOproblem

// package main

// import "fmt"

// type payment struct {
// 	getway stripepay
// }

// func (p payment) makePayment(amount float64) {
// 	// razorpaymentGw := razorpay{}
// 	// razorpaymentGw.pay(amount)
// 	// stripepaymentGw :=stripepay{}
// 	// stripepaymentGw.pay(amount)
// 	p.getway.pay(100)
// }

// type razorpay struct{}

// func (r razorpay) pay(amount float64) {
// 	fmt.Println("making payment with razorpay", amount)
// }

// type stripepay struct{}

// func (s stripepay) pay(amount float64) {
// 	fmt.Println("making payment with stripe", amount)
// }
// func main() {
// 	stripepaymentGw := stripepay{}
// 	newPayment := payment{getway: stripepaymentGw}
// 	newPayment.makePayment(100)
// }
//TODO solve

package main

import "fmt"

type paymenter interface {
	pay(amount float32)
}
type payment struct {
	gateway paymenter
}

func (p *payment) makePayment(amount float32) {
	p.gateway.pay(amount)
}

type razorpay struct{}

func (r razorpay) pay(amount float32) {
	fmt.Println("making payment with razorpay", amount)
}

type stripepay struct{}

func (s stripepay) pay(amount float32) {
	fmt.Println("making payment with stripepay", amount)
}

func main() {
	// stripepaymentGw := stripepay{}
	razorpaymentGw := razorpay{}
	newPayment := payment{gateway: razorpaymentGw}
	newPayment.makePayment(100)
}
