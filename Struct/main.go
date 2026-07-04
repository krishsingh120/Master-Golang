package main

import (
	"fmt"
	"time"
)

/*

- In go classes and objects and not available.
- In go structs and used.
- all type of work in go we use structs instead of using classes and objects.
- In go also create multiple instances of single struct.


default values : means by default assign all falsy values.
 - string ""
 - int 0
 - float 0
 - bool false

 - how to use constructor in go


*/

// syntax of structs => type <name_of_struct> struct {}
type Order struct {
	// all these are properties of struct
	id        string
	name      string
	price     float32
	status    string
	createdAt time.Time
}

// constructor type implementation in go.
func newOrder(id string, name string, price float32, status string, createdAt time.Time) *Order {
	order := Order{
		id:        id,
		name:      name,
		price:     price,
		status:    status,
		createdAt: createdAt,
	}

	return &order
}

// In struct we add a func they can use order properties.
// To use these properties we make a methods in go, these methods are called as behavior of the struct.
// syntax -> func receiver(w Worker) <Method_name> {}
// there are two types of receivers -> 1. value receivers 2. pointer receivers

// 1. value receiver
func (o Order) checkStatus() {
	fmt.Println(o.id)
	fmt.Println(o.name)
	fmt.Println(o.price)
	fmt.Println(o.status)
	fmt.Println(o.createdAt)
}

// 1. value receiver
func (o *Order) updateStatus(updatedStatus string) {
	o.status = updatedStatus
	fmt.Println(o.status)
}

// inheritance in go using embeddings
// Base struct
type Engine struct {
	Horsepower int
}

func (e Engine) Start() {
	fmt.Println("Vroom! Engine started.")
}

// Outer struct embedding 'Engine'
type Car struct {
	Engine // Anonymous field (Embedding)
	Model  string
}

func main() {
	// create instance of this order struct
	// order := Order{
	// 	id:        "123",
	// 	name:      "Iphone 17 Pro",
	// 	price:     1200.000,
	// 	status:    "PENDING",
	// 	createdAt: time.Now().UTC().Local(),
	// }

	// create order using constructor style
	order := newOrder("123", "Iphone 17 Pro", 1200.000, "PENDING", time.Now().UTC().Local())

	order.checkStatus() // call receiver by value
	order.updateStatus("COMPLETED")

	fmt.Println(order)

	// create struct without name and instantly assign value.
	language := struct {
		name string
		rate int
	}{"Hindi", 4}

	fmt.Println(language)


	hondaCity := Car{
		Engine: Engine{Horsepower: 12000},
		Model: "2026",
	}

	fmt.Println(hondaCity.Horsepower)
}
