package main

import (
	"fmt"
	"sync"
)

func multWorker(ch <-chan int, doubleCh chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()

	for value := range ch {
		doubleCh <- value * 2
	}
	close(doubleCh)
}

func printWorker(doubleCh <-chan int, wg *sync.WaitGroup) {
	defer wg.Done()

	for value := range doubleCh {
		fmt.Println(value)
	}
}

func main() {
	arr := []int{1, 2, 3, 4, 5}

	numCh := make(chan int, len(arr))
	doubleCh := make(chan int, len(arr))
	wg := sync.WaitGroup{}

	wg.Add(2)
	go multWorker(numCh, doubleCh, &wg)
	go printWorker(doubleCh, &wg)

	for _, value := range arr {
		numCh <- value
	}
	close(numCh)

	wg.Wait()
}
