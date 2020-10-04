package main

import (
	"fmt"
	"time"
)

func main() {
	o1 := order{id: 1, price: 1, product: "a", status: false}
	o2 := order{id: 2, price: 2, product: "b", status: false}
	o3 := order{id: 3, price: 3, product: "c", status: false}
	o4 := order{id: 4, price: 4, product: "d", status: false}
	o5 := order{id: 5, price: 5, product: "e", status: false}
	o6 := order{id: 6, price: 6, product: "f", status: false}
	o7 := order{id: 7, price: 7, product: "g", status: false}
	o8 := order{id: 8, price: 8, product: "h", status: false}
	o9 := order{id: 9, price: 9, product: "i", status: false}
	o10 := order{id: 10, price: 10, product: "j", status: false}
	o11 := order{id: 11, price: 11, product: "k", status: false}
	o12 := order{id: 12, price: 12, product: "l", status: false}

	o := orders{items: []order{o1, o2, o3, o4, o5, o6, o7, o8, o9, o10, o11, o12}}

	start := time.Now()
	c := o.process(4)
	for i := 0; i < len(o.items); i++ {
		o.items[i].status = true
		fmt.Println(<-c)
	}
	elapsed := time.Since(start)
	fmt.Println("elapsed:", elapsed)
}

type order struct {
	id      int
	price   int
	product string
	status  bool
}

type orders struct {
	items []order
}

func (o *orders) process(machines int) <-chan string {
	c := make(chan string)
	stride := len(o.items) / machines

	for g := 0; g < machines; g++ {
		go func(g int) {
			start := g * stride
			end := start + stride

			for _, v := range o.items[start:end] {
				prepareIngredients()
				grindCoffee()
				boilCoffee()
				c <- fmt.Sprintf("Machine: %d, Order: %s", g, v.product)
			}
		}(g)
	}

	return c
}

func prepareIngredients() {
	time.Sleep(time.Second / 2)
}

func grindCoffee() {
	time.Sleep(time.Second / 2)
}

func boilCoffee() {
	time.Sleep(time.Second)
}
