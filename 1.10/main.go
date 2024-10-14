package main

import "fmt"

func main() {
	arr := []float32{-25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5}

	tempGroups := make(map[int][]float32)

	for _, value := range arr {
		group := int(value) / 10 * 10

		tempGroups[group] = append(tempGroups[group], value)
	}

	fmt.Println(tempGroups)
}
