package main

import "fmt"

func main() {
	set1 := []int{1, 2, 3, 4, 5}
	set2 := []int{3, 4, 5, 6, 7}
	m := make(map[int]struct{})
	var res []int

	for _, value := range set1 {
		m[value] = struct{}{}
	}

	for _, value := range set2 {
		_, ok := m[value]
		if ok {
			res = append(res, value)
		}
	}

	fmt.Println(res)
}
