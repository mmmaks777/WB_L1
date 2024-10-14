package main

import (
	"fmt"
	"sync"
)

func main() {
	arr := []int{2, 4, 6, 8, 10}
	var sum int
	var wg sync.WaitGroup
	var mu sync.Mutex

	for _, value := range arr {
		wg.Add(1)
		go func() {
			defer wg.Done()
			mu.Lock()
			sum += value * value
			mu.Unlock()
		}()
	}
	wg.Wait()
	fmt.Println(sum)
}
