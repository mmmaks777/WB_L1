package main

import "fmt"

var justString string

func createHugeString(size int) string {
	bytes := make([]byte, size)
	for i := 0; i < size; i++ {
		bytes[i] = 'A'
	}
	return string(bytes)
}

func someFunc() {
	v := createHugeString(1 << 10)
	justString = string([]rune(v[:100]))
	fmt.Println(justString)
}

func main() {
	someFunc()
}
