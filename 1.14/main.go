package main

import "fmt"

func typeOf(i interface{}) {
	switch i.(type) {
	case int:
		fmt.Println("int type")
	case string:
		fmt.Println("string type")
	case bool:
		fmt.Println("bool type")
	case chan int:
		fmt.Println("chan int type")
	default:
		fmt.Println("unexpected type")
	}
}

func main() {
	x := 1
	typeOf(x)
}
