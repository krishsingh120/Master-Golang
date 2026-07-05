package main

import "fmt"

/*

Enums: Enumerated types

Enums: status = ["PENDING", "CONFIRM", "IN_PROGRESS"]
- In go Enums types not exist.
- we implement using consts.
- In go custom types also available.


*/

// custom types
// type orderStatus int

// const (
// 	Received orderStatus = iota + 1
// 	Confirmed
// 	Delivered
// 	InProgress
// )

type orderStatus string

const (
	Received   orderStatus = "Received"
	Confirmed  orderStatus = "Confirmed"
	Delivered  orderStatus = "Delivered"
	InProgress orderStatus = "InProgress"
)

func changeOrderStatus(status orderStatus) {
	fmt.Println("Current order status is:", status)
}

func main() {

	changeOrderStatus(InProgress)

}
