package main

import (
	"fmt"
)

// better solution
func findMissingNumber2(numbersArray []int) []int {

	result := make([]int, 0, len(numbersArray))
	count := make([]bool, len(numbersArray)+2)

	for _, i := range numbersArray {
		count[i-1] = true
	}

	for i := 0; i < len(count); i++ {
		if count[i] == false {
			result = append(result, i+1)
		}
	}

	return result
}

func main() {

	result := findMissingNumber2([]int{2, 3, 4, 5, 6})
	fmt.Println(result)
}
