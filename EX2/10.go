package main

import "fmt"

type any func(int) int

func Map(fn any, lists []int) []int {
	result := make([]int, 0, len(lists))
	for _, value := range lists {
		result = append(result, fn(value))
	}

	return result
}

func multiply(num int) int {
	return num * num
}

func main() {
	lists := []int{1, 2, 3, 4, 5}
	result := Map(multiply, lists)
	fmt.Println(result)
}
