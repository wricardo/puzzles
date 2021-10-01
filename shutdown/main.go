package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	done := make(chan struct{}, 0)
	wg := sync.WaitGroup{}
	wg.Add(1)
	go task1(done, &wg)
	close(done)
	wg.Wait()
}

func task1(done chan struct{}, wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		select {
		case <-done:
			return
		default:
			fmt.Println("perform task 1")
		}
		time.Sleep(time.Minute)
	}
}
