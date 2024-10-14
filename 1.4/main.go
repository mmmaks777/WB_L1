package main

import (
	"context"
	"flag"
	"fmt"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

func worker(ctx context.Context, ch chan int, num int, wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		select {
		case <-ctx.Done():
			fmt.Printf("worker %d is done\n", num)
			return
		case value, ok := <-ch:
			if !ok {
				fmt.Printf("worker %d is done - channel is closed\n", num)
				return
			}
			fmt.Printf("worker %d get a value: %d\n", num, value)
		}
	}
}

func main() {
	n := flag.Int("n", 0, "count of workers")
	flag.Parse()

	ch := make(chan int)
	ctx, cancel := context.WithCancel(context.Background())
	var wg sync.WaitGroup

	for i := 0; i < *n; i++ {
		wg.Add(1)
		go worker(ctx, ch, i, &wg)
	}

	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, syscall.SIGINT)
	go func() {
		<-sigCh
		fmt.Println("Cancel signal is recieved")
		cancel()
	}()

	i := 0
loop:
	for {
		select {
		case <-ctx.Done():
			close(ch)
			break loop
		default:
			ch <- i
			i++
			time.Sleep(time.Second)
		}
	}
	wg.Wait()
	fmt.Println("All workers are done")
}
