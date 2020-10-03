package main

import (
	"fmt"
	"sync"
	"time"
)

func findConcurrent(goroutines int, docs []string) {
	ch := make(chan string, len(docs))
	for _, doc := range docs {
		ch <- doc
	}
	close(ch)

	var wg sync.WaitGroup
	wg.Add(goroutines)

	for g := 0; g < goroutines; g++ {
		go func() {
			for d := range ch {
				read(time.Second * 5)
				fmt.Println(d)
			}
			wg.Done()
		}()
	}
	wg.Wait()
}

func read(seconds time.Duration) {
	time.Sleep(seconds)
}
