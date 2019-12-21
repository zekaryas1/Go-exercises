package main

import (
	"fmt"
)

type node struct {
	Value string
	Next  *node
}

type WordStack struct {
	Top     *node
	Size    int
	maxSize int
}

func (stack *WordStack) initialize(maxSize int) {
	stack.maxSize = maxSize
}

func (stack *WordStack) Push(value string) {
	if stack.Size >= stack.maxSize {
		fmt.Println("WordStack overflow")
		return
	}
	stack.Top = &node{
		Value: value,
		Next:  stack.Top,
	}
	stack.Size++
}

func (stack *WordStack) Pop() (value string) {
	if stack.Size > 0 {
		value = stack.Top.Value
		stack.Top = stack.Top.Next
		stack.Size--
		return
	}
	println("underflow")
	return ""
}
