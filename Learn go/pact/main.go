package main

import "fmt"

// func add(a, b int) int {
// 	return a + b
// }

// func addn(a ...int) (int,int) {
// 	sum := 0
// 	s:=len(a)

// 	for _, v := range a {
// 		sum += v
// 	}
// 	return sum,s
// }

// func counter(fn func(a int) int) int {
//  return fn(1);
// }

//	func counter() func() int {
//		co := 0
//		return func() int {
//			co++
//			return co
//		}
//	}
//
//	func changeNumber(a *int){
//	  *a = 5
//	  fmt.Println("change number in changeNumber function", *a)
//	}
// type user struct {
// 	name     string
// 	location string
// }
// type order struct {
// 	id int
// 	user
// 	status    string
// 	createdAt time.Time
// }

// func (o *order) changeOrderStatus(status string) {
// 	o.status = status
// }

// func new(id int, user user, status string) order {
// 	or := order{
// 		id:        id,
// 		user:      user,
// 		status:    status,
// 		createdAt: time.Now(),
// 	}
// 	return or
// }

// type paymenter interface{
// 	pay(amount int)
// }
// type payment struct {
//     gateway paymenter
// }

// func (p *payment) makePayment(amount int) {
// //   razorPaymentgw := razorPay{}
// //   razorPaymentgw.pay(amount)
// // stripePayment := stripe{}
// // stripePayment.pay(amount)
// p.gateway.pay(amount)
// }
// type stripe struct {}
// func (s *stripe) pay(amount int) {
// 	fmt.Println("payment made using stripe", amount)
// }
// type razorPay struct {}
// func (r *razorPay) pay(amount int) {
//  fmt.Println("payment made using razorpay", amount)
// }

// type OrderStatus int
type OrderStatus string

// const (
// 	Received OrderStatus = iota
// 	Confirmed
// 	Prepared
// 	Shipped
// 	Delivered
// )

// const (
// 	Received  OrderStatus = "Received"
// 	Confirmed             = "Confirmed"
// 	Prepared              = "Prepared"
// 	Shipped               = "Shipped"
// 	Delivered             = "Delivered"
// )

// func changeOrderStatus(order OrderStatus) {
// 	fmt.Println("order status changed to", order)
// }
 
// func printslice[T any](items []T) { T can be int | string | float | bool | comparable | any | interface{}
// 	for _, item := range items {
// 		fmt.Println(item)
// 	}
// }

type stack[T any] struct {
	element []T
}
func main() {
	// x := [2][2]int{{1, 2}, {3, 4}} // jodi array size na dei tahole seta k slice
	// xs := make([]int, 0, 5)
	// xs = append(xs, 1)
	// xs = append(xs, 200000)
	// // for i:=0;i<len(xs);i++{
	// //     fmt.Println(xs[i])
	// // }
	// for k, num := range xs {
	// 	fmt.Println(num, k)
	// }

	// xs = append(xs, 1)
	// xs = append(xs, 2)
	// xs2 := make([]int, len(xs))
	// copy(xs2, xs)
	// fmt.Println(slices.Equal(xs, xs2))
	// fmt.Println(xs2[0:2])

	// //map
	// m2 := map[string]string{"price": "20"}
	// k, ok := m2["price"]
	// if ok {
	// 	fmt.Println(k)
	// }
	// fmt.Println(m2)

	// m := make(map[string]string)
	// m["name"] = "John"
	// m["age"] = "30"
	// delete(m, "age")
	// clear(m)
	// fmt.Println(m)
	// arr := []int{1, 2, 3, 4, 5}
	// l,d := addn(arr...)
	// fmt.Println(d,l)

	// fn := func(a int) int {
	// 	return a + 4
	// }
	// fmt.Println(countp(fn))
	// c := counter()
	// fmt.Println(c())
	// fmt.Println(c())
	// fmt.Println(c())
	// fmt.Println(c())

	// a:=1
	// changeNumber(&a)
	// fmt.Println("change number of a in main function",a)
	// u := user{
	// 	name:     "majed",
	// 	location: "dhaka",
	// }
	// or := order{
	// 	id:        1,
	// 	user:      u,
	// 	status:    "pending",
	// 	createdAt: time.Now(),

	// oe:=new(1,u,"pending")
	// // }
	// oe.changeOrderStatus("completed")
	// fmt.Println(oe)

	// language := struct {
	// 	name   string
	// 	isGood bool
	// }{
	// 	"english", false,
	// }
	// fmt.Println(language)
	// stripePayment := stripe{}
	// // razorPaymentgw := razorPay{}
	// myPayment := payment{&stripePayment}
	// myPayment.makePayment(20)

	// changeOrderStatus(Received)
	// items := []string{"apple", "banana", "cherry"}
	// // items := []int{1, 2, 3, 4, 5}

	// // printslice(items)

	// myStack := stack[string]{items}
	// fmt.Printf("Stack: %v\n", myStack)



}
