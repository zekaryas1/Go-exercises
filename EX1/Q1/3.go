package main

import "fmt"

func main() {
	array := make([]int, 0, 10)
	for i := 0; i < 10; i++ {
		array = append(array, i)
	}
	fmt.Println(array)
}
