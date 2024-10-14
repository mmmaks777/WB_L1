package main

import "fmt"

func binarySearch(arr []int, target int) int {
	left, right := 0, len(arr)-1

	for left <= right {
		mid := left + (right-left)/2

		if arr[mid] == target {
			return mid
		} else if arr[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	return -1
}

func main() {
	arr := []int{1, 2, 4, 5, 6, 8, 11, 14, 15, 16}
	target := 11

	idx := binarySearch(arr, target)

	if idx != -1 {
		fmt.Println("Result: ", idx)
	} else {
		fmt.Println("Target not found")
	}
}
