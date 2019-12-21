package main

import "fmt"

func findMin(slicedArray []int) int {
	min := slicedArray[0]
	for _, value := range slicedArray {
		if value < min {
			min = value
		}
	}
	return min
}

func main() {
	values := []int{1, 2, -3, 14, 5, -6}
	maxElement := findMin(values)
	fmt.Println(maxElement)
}
