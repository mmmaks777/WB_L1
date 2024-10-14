package main

import (
	"time"
)

func Sleep(seconds int) {
	start := time.Now().Second()
	for {
		if time.Now().Second()-start >= seconds {
			return
		}
	}
}

func main() {
	Sleep(3)
}
