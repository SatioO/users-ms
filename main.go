package main

import (
	"fmt"
	"sync"
)

func main() {
	todos := Todos{}
	var wg sync.WaitGroup
	var mt sync.Mutex
	wg.Add(3)

	go func() {
		// Add Todo
		mt.Lock()
		todos.add(&Todo{Name: "Study Go", Status: false})
		mt.Unlock()
		wg.Done()
	}()

	go func() {
		// Add Todo
		mt.Lock()
		todos.add(&Todo{Name: "Study Serverless", Status: false})
		mt.Unlock()
		wg.Done()
	}()

	go func() {
		// Add Todo
		mt.Lock()
		todos.add(&Todo{Name: "Study Flutter", Status: false})
		mt.Unlock()
		wg.Done()
	}()

	wg.Wait()
	fmt.Println(todos.Items)
}
