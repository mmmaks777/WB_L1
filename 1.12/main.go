package main

import "fmt"

func main() {
	arr := []string{"cat", "cat", "dog", "cat", "tree"}
	res := make(map[string]struct{})

	for _, value := range arr {
		res[value] = struct{}{}
	}

	for key := range res {
		fmt.Println(key)
	}
}
