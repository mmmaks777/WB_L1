package main

import (
	"fmt"
	"math/big"
)

func main() {
	a := new(big.Int)
	b := new(big.Int)

	a.SetString("21048577", 10)
	b.SetString("2097153", 10)

	fmt.Println(new(big.Int).Add(a, b).String())

	fmt.Println(new(big.Int).Sub(a, b).String())

	fmt.Println(new(big.Int).Mul(a, b).String())

	fmt.Println(new(big.Int).Div(a, b).String())
}
