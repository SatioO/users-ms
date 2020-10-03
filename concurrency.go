package main

import (
	"fmt"
	"sync"
	"sync/atomic"
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

func addConcurrent(goroutines int, numbers []int) {
	var v int64
	totalNumbers := len(numbers)
	lastGoroutine := goroutines - 1
	stride := totalNumbers / goroutines
	var wg sync.WaitGroup
	wg.Add(goroutines)

	for g := 0; g < goroutines; g++ {
		go func(g int) {
			start := g * stride
			end := start + stride
			if g == lastGoroutine {
				end = totalNumbers
			}

			var lv int
			for _, n := range numbers[start:end] {
				lv += n
			}
			atomic.AddInt64(&v, int64(lv))
			wg.Done()
		}(g)
	}

	wg.Wait()

	fmt.Println(v)
}

func add(numbers []int) int {
	var sum int

	for _, v := range numbers {
		sum += v
	}

	return sum
}
