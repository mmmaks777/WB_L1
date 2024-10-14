package main

import (
	"flag"
	"fmt"
)

func main() {
	i := flag.Int("i", 0, "shift magnitude")
	setBit := flag.Bool("set", false, "Set the bit to 1 if true, else set to 0")
	flag.Parse()

	var num int64 = 4

	fmt.Printf("%064b\n", num)

	if *setBit == true {
		num |= 1 << *i
	} else {
		num &^= 1 << *i
	}

	fmt.Printf("%064b\n", num)
	fmt.Println(num)
}
