package main

import "fmt"

func printEr(nums ...int) {
	for _, value := range nums {
		fmt.Println(value)
	}
}

func main() {
	printEr(1, 2, 3)
	printEr(1, 2, 3, 4, )
	printEr(1, 2, 3, 4, 5)
}
