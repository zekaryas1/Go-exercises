package main

import "fmt"

func main() {
	for i := 0; i < 10; i++ {
		fmt.Printf("%v\n", i)
	}

	//scope error i isn't available here
	//fmt.Printf("%v\n", i)
}
