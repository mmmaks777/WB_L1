package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

type counter struct {
	cnt int32
}

func incWorker(c *counter, count int32, wg *sync.WaitGroup) {
	defer wg.Done()

	atomic.AddInt32(&c.cnt, count)
}

func main() {
	c := counter{}
	wg := sync.WaitGroup{}

	wg.Add(2)
	go incWorker(&c, 5, &wg)
	go incWorker(&c, 4, &wg)

	wg.Wait()

	fmt.Println(c.cnt)
}
