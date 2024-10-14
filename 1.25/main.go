package main

import (
	"time"
)

func Sleep(seconds int) {
	start := time.Now()

	for {
		if time.Since(start).Seconds() >= float64(seconds) {
			return
		}
	}
}

func main() {
	Sleep(5)
}
