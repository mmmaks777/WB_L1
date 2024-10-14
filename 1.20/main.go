package main

import (
	"fmt"
	"strings"
)

func reverseWords(str string) string {
	words := strings.Fields(str)
	n := len(words)

	for i := 0; i < n/2; i++ {
		words[i], words[n-i-1] = words[n-i-1], words[i]
	}

	return strings.Join(words, " ")
}

func main() {
	str := "snow dog sun"

	fmt.Println(reverseWords(str))
}
