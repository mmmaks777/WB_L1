package main

import "fmt"

func removeElement(arr []int, idx int) []int {
	if idx < 0 && idx > len(arr)-1 {
		fmt.Println("Index out of range")
		return nil
	}
	return append(arr[:idx], arr[idx+1:]...)
}

func main() {
	arr := []int{1, 2, 3, 4, 5}

	fmt.Println(removeElement(arr, 2))
}
