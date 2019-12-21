package main

import "fmt"

func findMax(slicedArray []int) int {
	max := slicedArray[0]
	for _, value := range slicedArray {
		if value > max {
			max = value
		}
	}
	return max
}

func main() {
	values := []int{1, 2, 3, 14, 5, 6}
	maxElement := findMax(values)
	fmt.Println(maxElement)
}
