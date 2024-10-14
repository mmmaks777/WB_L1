package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

var (
	stop bool
	mu   sync.RWMutex
)

func worker1(ch chan struct{}, wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		_, ok := <-ch
		if !ok {
			fmt.Println("worker1 is done")
			return
		}
	}
}

func worker2(ctx context.Context, wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		select {
		case <-ctx.Done():
			fmt.Println("worker2 is done")
			return
		}
	}
}

func worker3(ctx context.Context, wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		select {
		case <-ctx.Done():
			fmt.Println("worker3 is done")
			return
		}
	}
}

func worker4(wg *sync.WaitGroup) {
	defer wg.Done()

	for {
		mu.RLock()
		if stop {
			fmt.Println("worker4 is done")
			return
		}
		mu.RUnlock()
	}
}

func main() {
	wg := sync.WaitGroup{}
	// Способ 1 Quit Channel
	quitCh := make(chan struct{})

	wg.Add(1)
	go worker1(quitCh, &wg)
	close(quitCh)

	// Способ 2 Context with cancel
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	wg.Add(1)
	go worker2(ctx, &wg)
	cancel()

	// Способ 3 Context with timeout
	ctx, cancel = context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	wg.Add(1)
	go worker3(ctx, &wg)

	// Способ 4 Глобальная переменная-флаг с мьютексом
	wg.Add(1)
	go worker4(&wg)
	mu.Lock()
	stop = true
	mu.Unlock()

	wg.Wait()
}
