package main

import "fmt"

func plusTwo(x int) func(y int) int {

	return func(y int) int {
		return x + y
	}

}

func main() {
	p := plusTwo(12)
	fmt.Println(p(12))
}
