package main

import "fmt"

func reverseString(str string) string {
	runes := []rune(str)
	n := len(runes)

	for i := 0; i < n/2; i++ {
		runes[i], runes[n-i-1] = runes[n-i-1], runes[i]
	}

	return string(runes)
}

func main() {
	str := "главрыба"

	fmt.Println(reverseString(str))
}
