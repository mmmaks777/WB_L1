package main

import "fmt"

func quickSort(arr []int) {
	quickSortHelper(arr, 0, len(arr)-1)
}

func quickSortHelper(arr []int, low, high int) {
	if low < high {
		pivotIdx := medianOfThree(arr, low, high)

		pivotNewIdx := partition(arr, low, high, pivotIdx)

		quickSortHelper(arr, low, pivotNewIdx-1)
		quickSortHelper(arr, pivotNewIdx+1, high)
	}
}

func partition(arr []int, low, high, pivotIdx int) int {
	pivotValue := arr[pivotIdx]
	arr[pivotIdx], arr[high] = arr[high], arr[pivotIdx]

	storeIdx := low
	for i := low; i < high; i++ {
		if arr[i] < pivotValue {
			arr[i], arr[storeIdx] = arr[storeIdx], arr[i]
			storeIdx++
		}
	}

	arr[storeIdx], arr[high] = arr[high], arr[storeIdx]

	return storeIdx
}

func medianOfThree(arr []int, low, high int) int {
	mid := len(arr)/2 - 1
	if arr[low] > arr[high] {
		arr[low], arr[high] = arr[high], arr[low]
	}
	if arr[low] > arr[mid] {
		arr[low], arr[mid] = arr[mid], arr[low]
	}
	if arr[mid] > arr[high] {
		arr[mid], arr[high] = arr[high], arr[mid]
	}
	return mid
}

func main() {
	arr := []int{4, 5, 7, 8, 3, 9, 2, 10, 1, 6}

	quickSort(arr)

	fmt.Println(arr)
}
