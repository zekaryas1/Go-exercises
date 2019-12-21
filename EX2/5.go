package main

import "fmt"

func returnTwoParam(param1 int, param2 int) (int, int) {
	if param1 > param2 {
		return param2, param1
	}
	return param1, param2
}

func main() {
	fmt.Println(returnTwoParam(7, 2))
	fmt.Println(returnTwoParam(2, 7))
}
