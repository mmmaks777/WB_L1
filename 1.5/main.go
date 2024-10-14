package main

import (
	"context"
	"flag"
	"fmt"
	"sync"
	"time"
)

func worker(ch chan int, ctx context.Context, wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		select {
		case <-ctx.Done():
			fmt.Println("worker is done")
			return
		case value, ok := <-ch:
			if !ok {
				fmt.Println("channel is closed")
				return
			}
			fmt.Println(value)
		}
	}
}

func main() {
	n := flag.Int("n", 0, "time of life")
	flag.Parse()

	ch := make(chan int)
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(*n)*time.Second)
	defer cancel()
	wg := sync.WaitGroup{}

	wg.Add(1)
	go worker(ch, ctx, &wg)

	i := 1
loop:
	for {
		select {
		case <-ctx.Done():
			fmt.Println("Timeout")
			close(ch)
			break loop
		default:
			ch <- i
			i++
			time.Sleep(time.Second)
		}
	}
	wg.Wait()
}
