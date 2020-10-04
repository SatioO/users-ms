package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	c := fanIn(boring("John"), boring("Joe"))

	for i := 0; i < 10; i++ {
		fmt.Printf("You say: %q\n", <-c)
	}

	fmt.Println("Quit")
}

func fanIn(input1, input2 <-chan string) <-chan string {
	c := make(chan string)
	go func() {
		for {
			select {
			case s := <-input1:
				c <- s
			case s := <-input2:
				c <- s
			}
		}
	}()
	return c
}

func boring(name string) <-chan string {
	c := make(chan string)

	go func() {
		for i := 0; i < 5; i++ {
			c <- fmt.Sprintf("%s: %d", name, i)
			time.Sleep(time.Duration(rand.Intn(1e4)) * time.Millisecond)
		}
	}()

	return c
}
