package main

import (
	"fmt"
	"sync"
)

func main() {
	m := sync.Map{}
	wg := sync.WaitGroup{}

	wg.Add(5)
	for i := 0; i < 5; i++ {
		go func() {
			defer wg.Done()
			m.Store(i, fmt.Sprintf("worker %d", i))
		}()
	}
	wg.Wait()
	m.Range(func(key, value any) bool {
		fmt.Println(key, value)
		return true
	})
}
