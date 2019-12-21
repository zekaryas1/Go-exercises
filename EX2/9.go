package main

import "fmt"

func fib2(num int) {
	prev := 1
	next := 1
	oldNext := 0
	fmt.Printf("%d ", prev)
	fmt.Printf("%d ", next)
	for i := 1; i < num; i++ {
		oldNext = next
		next = next + prev
		fmt.Printf("%d ", next)
		prev = oldNext
	}
}

func main() {
	fib2(1000)
}
